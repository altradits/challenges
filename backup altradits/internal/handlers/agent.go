package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/altradits/altradits/internal/db"
	"github.com/altradits/altradits/internal/middleware"
	"github.com/altradits/altradits/internal/utils"
)

type AgentDashData struct {
	User               *middleware.User
	BalanceSats        int64
	BalanceKES         string
	KesPerSat          float64
	AgentFeePct        float64
	TxCountToday       int
	VolumeKESToday     string
	CommissionToday    int64
	CommissionKESToday string
	Transactions       []AgentTxRow
}

type AgentTxRow struct {
	TxRow
	Identifier string
}

func methodLabel(m string) string {
	if m == "mpesa" {
		return "M-Pesa"
	}
	return "cash"
}

func HandleAgentDashboard(w http.ResponseWriter, r *http.Request) {
	u := middleware.GetUser(r)

	var balanceSats int64
	db.DB.QueryRow("SELECT balance_sats FROM wallets WHERE user_id=$1", u.ID).Scan(&balanceSats)

	var kesPerSat, agentFeePct float64
	db.DB.QueryRow("SELECT kes_per_sat, agent_fee_pct FROM pool_settings LIMIT 1").Scan(&kesPerSat, &agentFeePct)

	var txCountToday int
	var volumeSatsToday int64
	db.DB.QueryRow(`SELECT COUNT(*), COALESCE(SUM(amount_sats),0) FROM transactions
		WHERE agent_id=$1 AND type IN ('agent_cashin','agent_cashout') AND created_at::date = CURRENT_DATE`, u.ID).
		Scan(&txCountToday, &volumeSatsToday)

	var commissionToday int64
	db.DB.QueryRow(`SELECT COALESCE(SUM(amount_sats),0) FROM transactions
		WHERE user_id=$1 AND type='commission' AND created_at::date = CURRENT_DATE`, u.ID).
		Scan(&commissionToday)

	rows, _ := db.DB.Query(`SELECT t.id, t.type, t.amount_sats, t.direction, t.status,
		COALESCE(t.reference,''), COALESCE(t.note,''), t.created_at, u2.identifier
		FROM transactions t JOIN users u2 ON u2.id = t.user_id
		WHERE t.user_id=$1 OR t.agent_id=$1
		ORDER BY t.created_at DESC LIMIT 20`, u.ID)
	var txs []AgentTxRow
	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			var tx AgentTxRow
			rows.Scan(&tx.ID, &tx.Type, &tx.AmountSats, &tx.Direction, &tx.Status,
				&tx.Reference, &tx.Note, &tx.CreatedAt, &tx.Identifier)
			tx.TimeAgo = utils.TimeAgo(tx.CreatedAt)
			txs = append(txs, tx)
		}
	}

	data := AgentDashData{
		User:               u,
		BalanceSats:        balanceSats,
		BalanceKES:         utils.FormatKES(utils.SatsToKES(balanceSats, kesPerSat)),
		KesPerSat:          kesPerSat,
		AgentFeePct:        agentFeePct * 100,
		TxCountToday:       txCountToday,
		VolumeKESToday:     utils.FormatKES(utils.SatsToKES(volumeSatsToday, kesPerSat)),
		CommissionToday:    commissionToday,
		CommissionKESToday: utils.FormatKES(utils.SatsToKES(commissionToday, kesPerSat)),
		Transactions:       txs,
	}
	serveTemplateData(w, r, "web/templates/agent/dashboard.html", data)
}

func HandleAgentCashIn(w http.ResponseWriter, r *http.Request) {
	u := middleware.GetUser(r)

	var kesPerSat, agentFeePct float64
	db.DB.QueryRow("SELECT kes_per_sat, agent_fee_pct FROM pool_settings LIMIT 1").Scan(&kesPerSat, &agentFeePct)

	data := map[string]interface{}{
		"User": u, "KesPerSat": kesPerSat, "AgentFeePct": agentFeePct * 100,
	}

	if r.Method == http.MethodGet {
		serveTemplateData(w, r, "web/templates/agent/cash_in.html", data)
		return
	}

	identifier := strings.TrimSpace(r.FormValue("identifier"))
	method := r.FormValue("method")
	reference := strings.TrimSpace(r.FormValue("reference"))
	kes, err := strconv.ParseFloat(r.FormValue("kes_amount"), 64)

	if identifier == "" || err != nil || kes <= 0 {
		data["Error"] = "Enter the customer's email/phone and a valid amount."
		serveTemplateData(w, r, "web/templates/agent/cash_in.html", data)
		return
	}

	var customerID string
	err = db.DB.QueryRow("SELECT id FROM users WHERE identifier=$1 AND role='customer'", identifier).Scan(&customerID)
	if err != nil {
		data["Error"] = "No Altradits customer found with that email or phone."
		serveTemplateData(w, r, "web/templates/agent/cash_in.html", data)
		return
	}

	commissionKES := kes * agentFeePct
	netKES := kes - commissionKES
	netSats := utils.KESToSats(netKES, kesPerSat)
	commissionSats := utils.KESToSats(commissionKES, kesPerSat)

	note := fmt.Sprintf("Cash-in via agent %s (%s)", u.Identifier, methodLabel(method))
	if reference != "" {
		note += " — ref " + reference
	}

	db.DB.Exec(`INSERT INTO transactions (user_id, type, amount_sats, direction, status, reference, note, agent_id, confirmed_at)
		VALUES ($1,'agent_cashin',$2,'in','confirmed',$3,$4,$5,NOW())`,
		customerID, netSats, reference, note, u.ID)
	db.DB.Exec("UPDATE wallets SET balance_sats=balance_sats+$1, updated_at=NOW() WHERE user_id=$2", netSats, customerID)

	if commissionSats > 0 {
		db.DB.Exec(`INSERT INTO transactions (user_id, type, amount_sats, direction, status, reference, note, agent_id, confirmed_at)
			VALUES ($1,'commission',$2,'in','confirmed',$3,$4,$1,NOW())`,
			u.ID, commissionSats, identifier, fmt.Sprintf("Commission — cash-in for %s", identifier))
		db.DB.Exec("UPDATE wallets SET balance_sats=balance_sats+$1, updated_at=NOW() WHERE user_id=$2", commissionSats, u.ID)
	}

	data["Success"] = fmt.Sprintf("Credited %s to %s (%s). Your commission: %s (%s).",
		utils.FormatSats(netSats), identifier, utils.FormatKES(netKES),
		utils.FormatSats(commissionSats), utils.FormatKES(commissionKES))
	serveTemplateData(w, r, "web/templates/agent/cash_in.html", data)
}

func HandleAgentCashOut(w http.ResponseWriter, r *http.Request) {
	u := middleware.GetUser(r)

	var kesPerSat, agentFeePct float64
	db.DB.QueryRow("SELECT kes_per_sat, agent_fee_pct FROM pool_settings LIMIT 1").Scan(&kesPerSat, &agentFeePct)

	data := map[string]interface{}{
		"User": u, "KesPerSat": kesPerSat, "AgentFeePct": agentFeePct * 100,
	}

	if r.Method == http.MethodGet {
		serveTemplateData(w, r, "web/templates/agent/cash_out.html", data)
		return
	}

	identifier := strings.TrimSpace(r.FormValue("identifier"))
	method := r.FormValue("method")
	reference := strings.TrimSpace(r.FormValue("reference"))
	sats, err := strconv.ParseInt(r.FormValue("sats_amount"), 10, 64)

	if identifier == "" || err != nil || sats <= 0 {
		data["Error"] = "Enter the customer's email/phone and a valid amount."
		serveTemplateData(w, r, "web/templates/agent/cash_out.html", data)
		return
	}

	var customerID string
	var balance int64
	err = db.DB.QueryRow(`SELECT u2.id, w.balance_sats FROM users u2
		JOIN wallets w ON w.user_id = u2.id
		WHERE u2.identifier=$1 AND u2.role='customer'`, identifier).Scan(&customerID, &balance)
	if err != nil {
		data["Error"] = "No Altradits customer found with that email or phone."
		serveTemplateData(w, r, "web/templates/agent/cash_out.html", data)
		return
	}
	if sats > balance {
		data["Error"] = fmt.Sprintf("Customer only has %s available.", utils.FormatSats(balance))
		serveTemplateData(w, r, "web/templates/agent/cash_out.html", data)
		return
	}

	commissionSats := int64(float64(sats) * agentFeePct)
	netSats := sats - commissionSats
	kesToGive := utils.SatsToKES(netSats, kesPerSat)
	commissionKES := utils.SatsToKES(commissionSats, kesPerSat)

	note := fmt.Sprintf("Cash-out via agent %s (%s)", u.Identifier, methodLabel(method))
	if reference != "" {
		note += " — ref " + reference
	}

	db.DB.Exec(`INSERT INTO transactions (user_id, type, amount_sats, direction, status, reference, note, agent_id, confirmed_at)
		VALUES ($1,'agent_cashout',$2,'out','confirmed',$3,$4,$5,NOW())`,
		customerID, sats, reference, note, u.ID)
	db.DB.Exec("UPDATE wallets SET balance_sats=balance_sats-$1, updated_at=NOW() WHERE user_id=$2 AND balance_sats>=$1", sats, customerID)

	if commissionSats > 0 {
		db.DB.Exec(`INSERT INTO transactions (user_id, type, amount_sats, direction, status, reference, note, agent_id, confirmed_at)
			VALUES ($1,'commission',$2,'in','confirmed',$3,$4,$1,NOW())`,
			u.ID, commissionSats, identifier, fmt.Sprintf("Commission — cash-out for %s", identifier))
		db.DB.Exec("UPDATE wallets SET balance_sats=balance_sats+$1, updated_at=NOW() WHERE user_id=$2", commissionSats, u.ID)
	}

	data["Success"] = fmt.Sprintf("Hand the customer %s via %s. Your commission: %s (%s).",
		utils.FormatKES(kesToGive), methodLabel(method),
		utils.FormatSats(commissionSats), utils.FormatKES(commissionKES))
	serveTemplateData(w, r, "web/templates/agent/cash_out.html", data)
}

func HandleAgentTransactions(w http.ResponseWriter, r *http.Request) {
	u := middleware.GetUser(r)

	rows, _ := db.DB.Query(`SELECT t.id, t.type, t.amount_sats, t.direction, t.status,
		COALESCE(t.reference,''), COALESCE(t.note,''), t.created_at, u2.identifier
		FROM transactions t JOIN users u2 ON u2.id = t.user_id
		WHERE t.user_id=$1 OR t.agent_id=$1
		ORDER BY t.created_at DESC LIMIT 100`, u.ID)
	var txs []AgentTxRow
	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			var tx AgentTxRow
			rows.Scan(&tx.ID, &tx.Type, &tx.AmountSats, &tx.Direction, &tx.Status,
				&tx.Reference, &tx.Note, &tx.CreatedAt, &tx.Identifier)
			tx.TimeAgo = utils.TimeAgo(tx.CreatedAt)
			txs = append(txs, tx)
		}
	}

	serveTemplateData(w, r, "web/templates/agent/transactions.html", map[string]interface{}{
		"User": u, "Transactions": txs,
	})
}
