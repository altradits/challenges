// Package pgdrv is a minimal PostgreSQL wire-protocol driver
// that registers itself as "postgres" with database/sql.
// It implements the simple query protocol sufficient for all Altradits queries.
package pgdrv

import (
	"context"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"database/sql/driver"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"net"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func init() {
	sql.Register("postgres", &Driver{})
}

// ---------------------------------------------------------------------------
// Driver
// ---------------------------------------------------------------------------

type Driver struct{}

func (d *Driver) Open(dsn string) (driver.Conn, error) {
	return open(dsn)
}

// ---------------------------------------------------------------------------
// Connection
// ---------------------------------------------------------------------------

type conn struct {
	nc      net.Conn
	pid     int32
	secret  int32
	txDepth int
}

func open(dsn string) (*conn, error) {
	cfg, err := parseDSN(dsn)
	if err != nil {
		return nil, err
	}
	addr := cfg.host + ":" + cfg.port
	nc, err := net.DialTimeout("tcp", addr, 10*time.Second)
	if err != nil {
		return nil, fmt.Errorf("pgdrv: dial %s: %w", addr, err)
	}
	c := &conn{nc: nc}
	if err := c.startup(cfg); err != nil {
		nc.Close()
		return nil, err
	}
	return c, nil
}

func (c *conn) Close() error { return c.nc.Close() }

func (c *conn) Begin() (driver.Tx, error) {
	_, err := c.exec("BEGIN")
	if err != nil {
		return nil, err
	}
	c.txDepth++
	return c, nil
}

func (c *conn) Commit() error {
	c.txDepth--
	_, err := c.exec("COMMIT")
	return err
}

func (c *conn) Rollback() error {
	c.txDepth--
	_, err := c.exec("ROLLBACK")
	return err
}

// ---------------------------------------------------------------------------
// Startup / auth
// ---------------------------------------------------------------------------

type config struct {
	host, port, user, password, dbname string
}

func parseDSN(dsn string) (config, error) {
	u, err := url.Parse(dsn)
	if err != nil {
		return config{}, err
	}
	host := u.Hostname()
	port := u.Port()
	if port == "" {
		port = "5432"
	}
	user := u.User.Username()
	pass, _ := u.User.Password()
	dbname := strings.TrimPrefix(u.Path, "/")
	return config{host: host, port: port, user: user, password: pass, dbname: dbname}, nil
}

func (c *conn) startup(cfg config) error {
	// StartupMessage
	params := map[string]string{
		"user":             cfg.user,
		"database":         cfg.dbname,
		"application_name": "altradits",
		"client_encoding":  "UTF8",
	}
	buf := []byte{0, 0, 0, 0, 0, 3, 0, 0} // length placeholder + protocol 3.0
	for k, v := range params {
		buf = append(buf, []byte(k)...)
		buf = append(buf, 0)
		buf = append(buf, []byte(v)...)
		buf = append(buf, 0)
	}
	buf = append(buf, 0)
	binary.BigEndian.PutUint32(buf[0:4], uint32(len(buf)))
	if _, err := c.nc.Write(buf); err != nil {
		return err
	}

	// Auth loop
	for {
		msg, body, err := c.readMsg()
		if err != nil {
			return err
		}
		switch msg {
		case 'R': // Authentication
			method := binary.BigEndian.Uint32(body[0:4])
			switch method {
			case 0: // OK
			case 5: // MD5
				salt := body[4:8]
				hash := md5pass(cfg.password, cfg.user, salt)
				if err := c.sendPasswordMsg("md5" + hash); err != nil {
					return err
				}
			case 10: // SASL (SCRAM-SHA-256)
				if err := c.authSCRAM(cfg, body[4:]); err != nil {
					return err
				}
			default:
				return fmt.Errorf("pgdrv: unsupported auth method %d", method)
			}
		case 'K': // BackendKeyData
			c.pid = int32(binary.BigEndian.Uint32(body[0:4]))
			c.secret = int32(binary.BigEndian.Uint32(body[4:8]))
		case 'Z': // ReadyForQuery
			return nil
		case 'E':
			return pgError(body)
		case 'S': // ParameterStatus — ignore
		default:
			return fmt.Errorf("pgdrv: unexpected startup msg %c", msg)
		}
	}
}

func md5pass(password, user string, salt []byte) string {
	h1 := md5.Sum([]byte(password + user))
	hex1 := hex.EncodeToString(h1[:])
	h2 := md5.Sum(append([]byte(hex1), salt...))
	return hex.EncodeToString(h2[:])
}

func (c *conn) sendPasswordMsg(pass string) error {
	buf := make([]byte, 5+len(pass)+1)
	buf[0] = 'p'
	binary.BigEndian.PutUint32(buf[1:5], uint32(4+len(pass)+1))
	copy(buf[5:], pass)
	_, err := c.nc.Write(buf)
	return err
}

// ---------------------------------------------------------------------------
// SCRAM-SHA-256 authentication (RFC 5802 / RFC 7677, no channel binding)
// ---------------------------------------------------------------------------

const gs2Header = "n,,"

func (c *conn) authSCRAM(cfg config, mechanisms []byte) error {
	if !hasSCRAMSHA256(mechanisms) {
		return fmt.Errorf("pgdrv: server does not offer SCRAM-SHA-256")
	}

	clientNonce := scramNonce()
	clientFirstBare := "n=,r=" + clientNonce
	if err := c.sendMsg('p', saslInitialResponse("SCRAM-SHA-256", gs2Header+clientFirstBare)); err != nil {
		return err
	}

	serverFirst, err := c.readSASLMsg(11)
	if err != nil {
		return err
	}
	attrs := parseSCRAMAttrs(serverFirst)
	serverNonce := attrs["r"]
	if !strings.HasPrefix(serverNonce, clientNonce) {
		return fmt.Errorf("pgdrv: SCRAM server nonce does not match client nonce")
	}
	salt, err := base64.StdEncoding.DecodeString(attrs["s"])
	if err != nil {
		return fmt.Errorf("pgdrv: bad SCRAM salt: %w", err)
	}
	iterations, err := strconv.Atoi(attrs["i"])
	if err != nil {
		return fmt.Errorf("pgdrv: bad SCRAM iteration count: %w", err)
	}

	// SASLprep is not applied: passwords are assumed to be plain ASCII.
	saltedPassword := pbkdf2SHA256([]byte(cfg.password), salt, iterations, sha256.Size)
	clientKey := hmacSHA256(saltedPassword, []byte("Client Key"))
	storedKey := sha256.Sum256(clientKey)

	clientFinalWithoutProof := "c=" + base64.StdEncoding.EncodeToString([]byte(gs2Header)) + ",r=" + serverNonce
	authMessage := clientFirstBare + "," + serverFirst + "," + clientFinalWithoutProof

	clientSignature := hmacSHA256(storedKey[:], []byte(authMessage))
	clientProof := xorBytes(clientKey, clientSignature)
	clientFinal := clientFinalWithoutProof + ",p=" + base64.StdEncoding.EncodeToString(clientProof)

	if err := c.sendMsg('p', []byte(clientFinal)); err != nil {
		return err
	}

	serverFinal, err := c.readSASLMsg(12)
	if err != nil {
		return err
	}
	finalAttrs := parseSCRAMAttrs(serverFinal)
	if errMsg, ok := finalAttrs["e"]; ok {
		return fmt.Errorf("pgdrv: SCRAM authentication failed: %s", errMsg)
	}
	serverSig, err := base64.StdEncoding.DecodeString(finalAttrs["v"])
	if err != nil {
		return fmt.Errorf("pgdrv: bad SCRAM server signature: %w", err)
	}
	serverKey := hmacSHA256(saltedPassword, []byte("Server Key"))
	if !hmac.Equal(serverSig, hmacSHA256(serverKey, []byte(authMessage))) {
		return fmt.Errorf("pgdrv: SCRAM server signature mismatch")
	}
	return nil
}

func hasSCRAMSHA256(mechanisms []byte) bool {
	for _, m := range splitNullTerminated(mechanisms) {
		if m == "SCRAM-SHA-256" {
			return true
		}
	}
	return false
}

func splitNullTerminated(b []byte) []string {
	var out []string
	start := 0
	for i, c := range b {
		if c == 0 {
			if i == start {
				break
			}
			out = append(out, string(b[start:i]))
			start = i + 1
		}
	}
	return out
}

func parseSCRAMAttrs(msg string) map[string]string {
	attrs := make(map[string]string)
	for _, part := range strings.Split(msg, ",") {
		if len(part) >= 2 && part[1] == '=' {
			attrs[part[:1]] = part[2:]
		}
	}
	return attrs
}

func scramNonce() string {
	b := make([]byte, 18)
	if _, err := rand.Read(b); err != nil {
		panic("pgdrv: crypto/rand unavailable: " + err.Error())
	}
	return base64.StdEncoding.EncodeToString(b)
}

func saslInitialResponse(mechanism, response string) []byte {
	buf := []byte(mechanism)
	buf = append(buf, 0)
	var l [4]byte
	binary.BigEndian.PutUint32(l[:], uint32(len(response)))
	buf = append(buf, l[:]...)
	buf = append(buf, response...)
	return buf
}

func (c *conn) sendMsg(msgType byte, body []byte) error {
	buf := make([]byte, 5+len(body))
	buf[0] = msgType
	binary.BigEndian.PutUint32(buf[1:5], uint32(4+len(body)))
	copy(buf[5:], body)
	_, err := c.nc.Write(buf)
	return err
}

func (c *conn) readSASLMsg(expectType uint32) (string, error) {
	msg, body, err := c.readMsg()
	if err != nil {
		return "", err
	}
	if msg != 'R' {
		return "", fmt.Errorf("pgdrv: expected authentication message, got %c", msg)
	}
	got := binary.BigEndian.Uint32(body[0:4])
	if got != expectType {
		return "", fmt.Errorf("pgdrv: expected SASL auth type %d, got %d", expectType, got)
	}
	return string(body[4:]), nil
}

func hmacSHA256(key, data []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}

func xorBytes(a, b []byte) []byte {
	out := make([]byte, len(a))
	for i := range a {
		out[i] = a[i] ^ b[i]
	}
	return out
}

// pbkdf2SHA256 implements PBKDF2 (RFC 2898) with HMAC-SHA256, used to derive
// SCRAM's SaltedPassword from the connection password.
func pbkdf2SHA256(password, salt []byte, iterations, keyLen int) []byte {
	prf := hmac.New(sha256.New, password)
	hashLen := prf.Size()
	numBlocks := (keyLen + hashLen - 1) / hashLen
	dk := make([]byte, 0, numBlocks*hashLen)
	for block := 1; block <= numBlocks; block++ {
		prf.Reset()
		prf.Write(salt)
		var be [4]byte
		binary.BigEndian.PutUint32(be[:], uint32(block))
		prf.Write(be[:])
		u := prf.Sum(nil)
		t := make([]byte, len(u))
		copy(t, u)
		for i := 1; i < iterations; i++ {
			prf.Reset()
			prf.Write(u)
			u = prf.Sum(nil)
			for j := range t {
				t[j] ^= u[j]
			}
		}
		dk = append(dk, t...)
	}
	return dk[:keyLen]
}

// ---------------------------------------------------------------------------
// Query execution (simple query protocol)
// ---------------------------------------------------------------------------

type stmt struct {
	c     *conn
	query string
}

func (c *conn) Prepare(query string) (driver.Stmt, error) {
	return &stmt{c: c, query: query}, nil
}

func (s *stmt) Close() error                                    { return nil }
func (s *stmt) NumInput() int                                   { return -1 }
func (s *stmt) Exec(args []driver.Value) (driver.Result, error) { return s.c.execStmt(s.query, args) }
func (s *stmt) Query(args []driver.Value) (driver.Rows, error)  { return s.c.queryStmt(s.query, args) }

// ExecerContext for context support
func (c *conn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	vals := namedToValues(args)
	return c.execStmt(query, vals)
}

func (c *conn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	vals := namedToValues(args)
	return c.queryStmt(query, vals)
}

func namedToValues(args []driver.NamedValue) []driver.Value {
	vals := make([]driver.Value, len(args))
	for i, a := range args {
		vals[i] = a.Value
	}
	return vals
}

func (c *conn) exec(q string) (driver.Result, error) {
	return c.execStmt(q, nil)
}

func (c *conn) execStmt(query string, args []driver.Value) (driver.Result, error) {
	q := interpolate(query, args)
	rows, err := c.simpleQuery(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// drain
	for {
		if err := rows.Next(nil); err == io.EOF {
			break
		} else if err != nil {
			return rows.result, err
		}
	}
	return rows.result, nil
}

func (c *conn) queryStmt(query string, args []driver.Value) (*rows, error) {
	q := interpolate(query, args)
	return c.simpleQuery(q)
}

// ---------------------------------------------------------------------------
// Simple query protocol
// ---------------------------------------------------------------------------

type result struct {
	rowsAffected int64
	lastInsertID int64
}

func (r result) LastInsertId() (int64, error) { return r.lastInsertID, nil }
func (r result) RowsAffected() (int64, error) { return r.rowsAffected, nil }

type rows struct {
	c       *conn
	cols    []string
	colOIDs []uint32
	data    [][]driver.Value
	pos     int
	result  result
	done    bool
	pending *pendingMsg
}

// pendingMsg holds a message read while resolving the query's response
// (e.g. CommandComplete for a statement with no RowDescription) so Next
// can process it on its first call instead of dropping it.
type pendingMsg struct {
	msg  byte
	body []byte
}

func (r *rows) Columns() []string { return r.cols }

func (r *rows) Close() error {
	if !r.done {
		// drain remaining
		for {
			if err := r.Next(nil); err == io.EOF {
				break
			}
		}
	}
	return nil
}

func (r *rows) Next(dest []driver.Value) error {
	if r.pos < len(r.data) {
		if dest != nil {
			row := r.data[r.pos]
			for i, v := range row {
				dest[i] = v
			}
		}
		r.pos++
		return nil
	}
	if r.done {
		return io.EOF
	}
	var msg byte
	var body []byte
	if r.pending != nil {
		msg, body = r.pending.msg, r.pending.body
		r.pending = nil
	} else {
		var err error
		msg, body, err = r.c.readMsg()
		if err != nil {
			r.done = true
			return err
		}
	}
	switch msg {
	case 'D': // DataRow
		numCols := int(binary.BigEndian.Uint16(body[0:2]))
		row := make([]driver.Value, numCols)
		off := 2
		for i := 0; i < numCols; i++ {
			length := int32(binary.BigEndian.Uint32(body[off:]))
			off += 4
			if length < 0 {
				row[i] = nil // NULL
				continue
			}
			row[i] = decodeValue(string(body[off:off+int(length)]), r.colOIDs[i])
			off += int(length)
		}
		r.data = append(r.data, row)
		return r.Next(dest)
	case 'C': // CommandComplete
		tag := strings.TrimRight(string(body), "\x00")
		parts := strings.Fields(tag)
		if len(parts) > 0 {
			n, _ := strconv.ParseInt(parts[len(parts)-1], 10, 64)
			r.result.rowsAffected = n
		}
		return r.Next(dest)
	case 'Z': // ReadyForQuery
		r.done = true
		return io.EOF
	case 'I': // EmptyQueryResponse
		r.done = true
		return io.EOF
	case 'E':
		r.done = true
		if err := r.c.drainToReady(); err != nil {
			return err
		}
		return pgError(body)
	case 'S': // ParameterStatus
		return r.Next(dest)
	default:
		return r.Next(dest)
	}
}

func (c *conn) simpleQuery(q string) (*rows, error) {
	// Q message
	buf := make([]byte, 5+len(q)+1)
	buf[0] = 'Q'
	binary.BigEndian.PutUint32(buf[1:5], uint32(4+len(q)+1))
	copy(buf[5:], q)
	if _, err := c.nc.Write(buf); err != nil {
		return nil, err
	}

	// Resolve the column set (if any) before returning, so Columns() is
	// accurate on the first call — database/sql sizes Next's dest from it.
	r := &rows{c: c}
	for {
		msg, body, err := c.readMsg()
		if err != nil {
			return nil, err
		}
		switch msg {
		case 'T': // RowDescription
			r.cols, r.colOIDs = parseRowDescription(body)
			return r, nil
		case 'S': // ParameterStatus
			continue
		case 'E':
			if err := c.drainToReady(); err != nil {
				return nil, err
			}
			return nil, pgError(body)
		default:
			// No RowDescription coming (e.g. CommandComplete for an
			// INSERT/UPDATE/DELETE, or EmptyQueryResponse/ReadyForQuery).
			r.pending = &pendingMsg{msg: msg, body: body}
			return r, nil
		}
	}
}

// drainToReady reads until ReadyForQuery, which Postgres always sends after
// an ErrorResponse, so the connection stays in sync for the next query.
func (c *conn) drainToReady() error {
	for {
		msg, _, err := c.readMsg()
		if err != nil {
			return err
		}
		if msg == 'Z' {
			return nil
		}
	}
}

// parseRowDescription returns each column's name and PostgreSQL data type
// OID. The 18 bytes of field metadata following the null-terminated name are:
// Int32 table OID, Int16 attnum, Int32 type OID, Int16 typlen, Int32 typmod,
// Int16 format code.
func parseRowDescription(body []byte) ([]string, []uint32) {
	numCols := int(binary.BigEndian.Uint16(body[0:2]))
	cols := make([]string, numCols)
	oids := make([]uint32, numCols)
	off := 2
	for i := 0; i < numCols; i++ {
		end := off
		for body[end] != 0 {
			end++
		}
		cols[i] = string(body[off:end])
		meta := end + 1
		// meta+0:4 = table OID, meta+4:6 = attnum, meta+6:10 = type OID
		oids[i] = binary.BigEndian.Uint32(body[meta+6 : meta+10])
		off = meta + 18 // skip field metadata (18 bytes)
	}
	return cols, oids
}

// PostgreSQL type OIDs for columns that should be decoded into time.Time.
const (
	oidDate        uint32 = 1082
	oidTimestamp   uint32 = 1114
	oidTimestamptz uint32 = 1184
)

// timestampLayouts are tried in order when decoding a date/timestamp/
// timestamptz column's text value. Postgres only includes fractional
// seconds and/or a UTC offset when non-zero, so several layouts are needed.
var timestampLayouts = []string{
	"2006-01-02 15:04:05.999999-07:00",
	"2006-01-02 15:04:05.999999-07",
	"2006-01-02 15:04:05-07:00",
	"2006-01-02 15:04:05-07",
	"2006-01-02 15:04:05.999999",
	"2006-01-02 15:04:05",
	"2006-01-02",
}

func parseTimestamp(s string) (time.Time, error) {
	for _, layout := range timestampLayouts {
		if t, err := time.Parse(layout, s); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("pgdrv: cannot parse timestamp %q", s)
}

// decodeValue converts a column's text value to the driver.Value type
// database/sql expects. date/timestamp/timestamptz columns are parsed into
// time.Time so they can be scanned directly into *time.Time fields;
// everything else is returned as a string (numeric/bool destinations are
// converted from string by database/sql itself).
func decodeValue(s string, oid uint32) driver.Value {
	switch oid {
	case oidDate, oidTimestamp, oidTimestamptz:
		if t, err := parseTimestamp(s); err == nil {
			return t
		}
	}
	return s
}

// ---------------------------------------------------------------------------
// Wire protocol helpers
// ---------------------------------------------------------------------------

func (c *conn) readMsg() (byte, []byte, error) {
	header := make([]byte, 5)
	if _, err := io.ReadFull(c.nc, header); err != nil {
		return 0, nil, fmt.Errorf("pgdrv: readMsg header: %w", err)
	}
	msgType := header[0]
	length := int(binary.BigEndian.Uint32(header[1:5])) - 4
	if length < 0 {
		length = 0
	}
	body := make([]byte, length)
	if length > 0 {
		if _, err := io.ReadFull(c.nc, body); err != nil {
			return 0, nil, fmt.Errorf("pgdrv: readMsg body: %w", err)
		}
	}
	return msgType, body, nil
}

func pgError(body []byte) error {
	fields := map[byte]string{}
	i := 0
	for i < len(body) {
		ft := body[i]
		if ft == 0 {
			break
		}
		i++
		end := i
		for end < len(body) && body[end] != 0 {
			end++
		}
		fields[ft] = string(body[i:end])
		i = end + 1
	}
	return fmt.Errorf("pgdrv: %s: %s", fields['C'], fields['M'])
}

// ---------------------------------------------------------------------------
// Value interpolation (converts $1, $2 placeholders to literal values)
// ---------------------------------------------------------------------------

func interpolate(query string, args []driver.Value) string {
	if len(args) == 0 {
		return query
	}
	for i, arg := range args {
		placeholder := fmt.Sprintf("$%d", i+1)
		query = strings.ReplaceAll(query, placeholder, sqlLiteral(arg))
	}
	return query
}

func sqlLiteral(v driver.Value) string {
	if v == nil {
		return "NULL"
	}
	switch val := v.(type) {
	case int64:
		return strconv.FormatInt(val, 10)
	case float64:
		if math.IsNaN(val) || math.IsInf(val, 0) {
			return "NULL"
		}
		return strconv.FormatFloat(val, 'f', -1, 64)
	case bool:
		if val {
			return "TRUE"
		}
		return "FALSE"
	case []byte:
		return "'" + escapeSingle(string(val)) + "'"
	case string:
		return "'" + escapeSingle(val) + "'"
	case time.Time:
		return "'" + val.UTC().Format(time.RFC3339Nano) + "'"
	default:
		return "'" + escapeSingle(fmt.Sprintf("%v", val)) + "'"
	}
}

func escapeSingle(s string) string {
	return strings.ReplaceAll(s, "'", "''")
}
