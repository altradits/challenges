package handlers

import (
	"net/http"
	"time"

	"github.com/altradits/altradits/internal/db"
	"github.com/altradits/altradits/internal/middleware"
	"github.com/altradits/altradits/internal/utils"
)

type AdminDashData struct {
	User           *middleware.User
	TotalCustomers int
	NewToday       int
	TotalAUMSats   int64
	TotalAUMBTC    string
	TotalAUMKES    string
	PendingDeposits int
	PendingWithdrawals int
	RecentTx       []TxRow
}

func HandleAdminDashboard(w http.ResponseWriter, r *http.Request) {
	u := middleware.GetUser(r)
	var d AdminDashData
	d.User = u

	var kesPerSat float64
	db.DB.QueryRow("SELECT kes_per_sat FROM pool_settings LIMIT 1").Scan(&kesPerSat)

	db.DB.QueryRow("SELECT COUNT(*) FROM users WHERE role='customer'").Scan(&d.TotalCustomers)
	db.DB.QueryRow("SELECT COUNT(*) FROM users WHERE role='customer' AND created_at::date = CURRENT_DATE").Scan(&d.NewToday)
	db.DB.QueryRow("SELECT COALESCE(SUM(balance_sats),0) FROM wallets").Scan(&d.TotalAUMSats)
	d.TotalAUMBTC = utils.FormatBTC(d.TotalAUMSats)
	d.TotalAUMKES = utils.FormatKESAmount(utils.SatsToKES(d.TotalAUMSats, kesPerSat))
	db.DB.QueryRow("SELECT COUNT(*) FROM transactions WHERE type='deposit' AND status='pending'").Scan(&d.PendingDeposits)
	db.DB.QueryRow("SELECT COUNT(*) FROM transactions WHERE type='withdrawal' AND status='pending'").Scan(&d.PendingWithdrawals)

	rows, _ := db.DB.Query(`SELECT t.id, t.type, t.amount_sats, t.direction, t.status,
		COALESCE(t.reference,''), COALESCE(t.note,''), t.created_at
		FROM transactions t ORDER BY t.created_at DESC LIMIT 20`)
	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			var tx TxRow
			rows.Scan(&tx.ID, &tx.Type, &tx.AmountSats, &tx.Direction,
				&tx.Status, &tx.Reference, &tx.Note, &tx.CreatedAt)
			tx.TimeAgo = utils.TimeAgo(tx.CreatedAt)
			tx.AmountBTC = utils.FormatBTC(tx.AmountSats)
			tx.AmountKES = utils.FormatKESAmount(utils.SatsToKES(tx.AmountSats, kesPerSat))
			d.RecentTx = append(d.RecentTx, tx)
		}
	}

	serveTemplateData(w, r, "web/templates/admin/dashboard.html", d)
}

type CustomerRow struct {
	ID          string
	Identifier  string
	Role        string
	FullName    string
	BalanceSats int64
	BalanceBTC  string
	BalanceKES  string
	LastSeen    string
	CreatedAt   time.Time
	IsBlocked   bool
}

func HandleAdminCustomers(w http.ResponseWriter, r *http.Request) {
	u := middleware.GetUser(r)
	search := r.URL.Query().Get("q")

	var kesPerSat float64
	db.DB.QueryRow("SELECT kes_per_sat FROM pool_settings LIMIT 1").Scan(&kesPerSat)

	query := `SELECT u.id, u.identifier, u.role, COALESCE(u.full_name,''),
		COALESCE(w.balance_sats,0), COALESCE(u.last_seen::text,'never'), u.created_at, u.is_blocked
		FROM users u LEFT JOIN wallets w ON w.user_id=u.id
		WHERE u.role IN ('customer','trader')`
	if search != "" {
		query += " AND u.identifier ILIKE '%" + search + "%'"
	}
	query += " ORDER BY u.created_at DESC LIMIT 100"

	rows, _ := db.DB.Query(query)
	var customers []CustomerRow
	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			var c CustomerRow
			var lastSeen string
			rows.Scan(&c.ID, &c.Identifier, &c.Role, &c.FullName,
				&c.BalanceSats, &lastSeen, &c.CreatedAt, &c.IsBlocked)
			c.LastSeen = lastSeen
			c.BalanceBTC = utils.FormatBTC(c.BalanceSats)
			c.BalanceKES = utils.FormatKESAmount(utils.SatsToKES(c.BalanceSats, kesPerSat))
			customers = append(customers, c)
		}
	}

	serveTemplateData(w, r, "web/templates/admin/customers.html", map[string]interface{}{
		"User": u, "Customers": customers, "Search": search,
	})
}

func HandleAdminDeposits(w http.ResponseWriter, r *http.Request) {
	u := middleware.GetUser(r)

	var kesPerSat float64
	db.DB.QueryRow("SELECT kes_per_sat FROM pool_settings LIMIT 1").Scan(&kesPerSat)

	if r.Method == http.MethodPost {
		txID := r.FormValue("tx_id")
		action := r.FormValue("action")
		if action == "approve" {
			// Get transaction details
			var userID string
			var sats int64
			db.DB.QueryRow("SELECT user_id, amount_sats FROM transactions WHERE id=$1", txID).Scan(&userID, &sats)
			// Credit wallet
			db.DB.Exec("UPDATE wallets SET balance_sats=balance_sats+$1, updated_at=NOW() WHERE user_id=$2", sats, userID)
			db.DB.Exec("UPDATE transactions SET status='confirmed', confirmed_at=NOW() WHERE id=$1", txID)
			db.DB.Exec("INSERT INTO audit_logs (actor_id, action, target_id, detail) VALUES ($1,'deposit_approved',$2,$3)",
				u.ID, txID, "approved")
		} else if action == "reject" {
			db.DB.Exec("UPDATE transactions SET status='failed' WHERE id=$1", txID)
		}
	}

	rows, _ := db.DB.Query(`SELECT t.id, t.type, t.amount_sats, t.direction, t.status,
		COALESCE(t.reference,''), COALESCE(t.note,''), t.created_at, u.identifier
		FROM transactions t JOIN users u ON u.id=t.user_id
		WHERE t.type='deposit' ORDER BY t.created_at DESC LIMIT 50`)
	var txs []struct {
		TxRow
		Identifier string
	}
	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			var tx struct {
				TxRow
				Identifier string
			}
			rows.Scan(&tx.ID, &tx.Type, &tx.AmountSats, &tx.Direction,
				&tx.Status, &tx.Reference, &tx.Note, &tx.CreatedAt, &tx.Identifier)
			tx.TimeAgo = utils.TimeAgo(tx.CreatedAt)
			tx.AmountBTC = utils.FormatBTC(tx.AmountSats)
			tx.AmountKES = utils.FormatKESAmount(utils.SatsToKES(tx.AmountSats, kesPerSat))
			txs = append(txs, tx)
		}
	}

	serveTemplateData(w, r, "web/templates/admin/deposits.html", map[string]interface{}{
		"User": u, "Transactions": txs,
	})
}

func HandleAdminWithdrawals(w http.ResponseWriter, r *http.Request) {
	u := middleware.GetUser(r)

	var kesPerSat float64
	db.DB.QueryRow("SELECT kes_per_sat FROM pool_settings LIMIT 1").Scan(&kesPerSat)

	if r.Method == http.MethodPost {
		txID := r.FormValue("tx_id")
		action := r.FormValue("action")
		if action == "approve" {
			var userID string
			var sats int64
			db.DB.QueryRow("SELECT user_id, amount_sats FROM transactions WHERE id=$1", txID).Scan(&userID, &sats)
			db.DB.Exec("UPDATE wallets SET balance_sats=balance_sats-$1, updated_at=NOW() WHERE user_id=$2 AND balance_sats>=$1", sats, userID)
			db.DB.Exec("UPDATE transactions SET status='confirmed', confirmed_at=NOW() WHERE id=$1", txID)
		} else if action == "reject" {
			db.DB.Exec("UPDATE transactions SET status='failed' WHERE id=$1", txID)
		}
	}

	rows, _ := db.DB.Query(`SELECT t.id, t.type, t.amount_sats, t.direction, t.status,
		COALESCE(t.reference,''), COALESCE(t.note,''), t.created_at, u.identifier
		FROM transactions t JOIN users u ON u.id=t.user_id
		WHERE t.type='withdrawal' ORDER BY t.created_at DESC LIMIT 50`)
	var txs []struct {
		TxRow
		Identifier string
	}
	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			var tx struct {
				TxRow
				Identifier string
			}
			rows.Scan(&tx.ID, &tx.Type, &tx.AmountSats, &tx.Direction,
				&tx.Status, &tx.Reference, &tx.Note, &tx.CreatedAt, &tx.Identifier)
			tx.TimeAgo = utils.TimeAgo(tx.CreatedAt)
			tx.AmountBTC = utils.FormatBTC(tx.AmountSats)
			tx.AmountKES = utils.FormatKESAmount(utils.SatsToKES(tx.AmountSats, kesPerSat))
			txs = append(txs, tx)
		}
	}

	serveTemplateData(w, r, "web/templates/admin/withdrawals.html", map[string]interface{}{
		"User": u, "Transactions": txs,
	})
}

func HandleAdminSettings(w http.ResponseWriter, r *http.Request) {
	u := middleware.GetUser(r)
	msg := ""

	if r.Method == http.MethodPost {
		kesPerSat := r.FormValue("kes_per_sat")
		targetYield := r.FormValue("target_yield")
		adminFee := r.FormValue("admin_fee")
		db.DB.Exec("UPDATE pool_settings SET kes_per_sat=$1, target_yield=$2, admin_fee_pct=$3, updated_at=NOW(), updated_by=$4",
			kesPerSat, targetYield, adminFee, u.ID)
		msg = "Settings updated."
	}

	var kesPerSat, targetYield, adminFee float64
	db.DB.QueryRow("SELECT kes_per_sat, target_yield, admin_fee_pct FROM pool_settings LIMIT 1").
		Scan(&kesPerSat, &targetYield, &adminFee)

	serveTemplateData(w, r, "web/templates/admin/settings.html", map[string]interface{}{
		"User": u, "KesPerSat": kesPerSat, "TargetYield": targetYield * 100,
		"AdminFee": adminFee * 100, "Message": msg,
	})
}

func HandleAdminBlockUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/admin/customers", http.StatusSeeOther)
		return
	}
	u := middleware.GetUser(r)
	targetID := r.FormValue("user_id")
	action := r.FormValue("action")
	blocked := action == "block"
	db.DB.Exec("UPDATE users SET is_blocked=$1 WHERE id=$2", blocked, targetID)
	db.DB.Exec("INSERT INTO audit_logs (actor_id, action, target_id) VALUES ($1,$2,$3)",
		u.ID, action+"_user", targetID)
	http.Redirect(w, r, "/admin/customers", http.StatusSeeOther)
}
