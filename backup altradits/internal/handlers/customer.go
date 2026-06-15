package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/altradits/altradits/internal/db"
	"github.com/altradits/altradits/internal/middleware"
	"github.com/altradits/altradits/internal/utils"
)

type WalletData struct {
	User              *middleware.User
	BalanceSats       int64
	BalanceKES        string
	LightningAddr     string
	InterestThisMonth int64
	Transactions      []TxRow
	KesPerSat         float64
}

type TxRow struct {
	ID          string
	Type        string
	AmountSats  int64
	AmountBTC   string
	AmountKES   string
	Direction   string
	Status      string
	Reference   string
	Note        string
	TimeAgo     string
	CreatedAt   time.Time
}

func HandleCustomerDashboard(w http.ResponseWriter, r *http.Request) {
	u := middleware.GetUser(r)

	var balanceSats int64
	var lightningAddr string
	db.DB.QueryRow(`SELECT balance_sats, COALESCE(lightning_addr,'') FROM wallets WHERE user_id=$1`, u.ID).
		Scan(&balanceSats, &lightningAddr)

	var kesPerSat float64
	db.DB.QueryRow("SELECT kes_per_sat FROM pool_settings LIMIT 1").Scan(&kesPerSat)

	var interestMonth int64
	db.DB.QueryRow(`SELECT COALESCE(SUM(amount_sats),0) FROM transactions
		WHERE user_id=$1 AND type='interest' AND created_at >= date_trunc('month', NOW())`, u.ID).
		Scan(&interestMonth)

	rows, _ := db.DB.Query(`SELECT id, type, amount_sats, direction, status,
		COALESCE(reference,''), COALESCE(note,''), created_at
		FROM transactions WHERE user_id=$1 ORDER BY created_at DESC LIMIT 10`, u.ID)
	var txs []TxRow
	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			var tx TxRow
			rows.Scan(&tx.ID, &tx.Type, &tx.AmountSats, &tx.Direction,
				&tx.Status, &tx.Reference, &tx.Note, &tx.CreatedAt)
			tx.TimeAgo = utils.TimeAgo(tx.CreatedAt)
			txs = append(txs, tx)
		}
	}

	data := WalletData{
		User:              u,
		BalanceSats:       balanceSats,
		BalanceKES:        utils.FormatKES(utils.SatsToKES(balanceSats, kesPerSat)),
		LightningAddr:     lightningAddr,
		InterestThisMonth: interestMonth,
		Transactions:      txs,
		KesPerSat:         kesPerSat,
	}
	serveTemplateData(w, r, "web/templates/customer/dashboard.html", data)
}

func HandleDeposit(w http.ResponseWriter, r *http.Request) {
	u := middleware.GetUser(r)
	var kesPerSat float64
	db.DB.QueryRow("SELECT kes_per_sat FROM pool_settings LIMIT 1").Scan(&kesPerSat)

	if r.Method == http.MethodGet {
		serveTemplateData(w, r, "web/templates/customer/deposit.html", map[string]interface{}{
			"User": u, "KesPerSat": kesPerSat,
		})
		return
	}

	kesStr := r.FormValue("kes_amount")
	kes, err := strconv.ParseFloat(kesStr, 64)
	if err != nil || kes <= 0 {
		serveTemplateData(w, r, "web/templates/customer/deposit.html", map[string]interface{}{
			"User": u, "KesPerSat": kesPerSat, "Error": "Enter a valid amount.",
		})
		return
	}

	sats := utils.KESToSats(kes, kesPerSat)
	mpesaRef := r.FormValue("mpesa_ref")
	if mpesaRef == "" {
		mpesaRef = "PENDING"
	}

	// Manual entry: create pending transaction
	db.DB.Exec(`INSERT INTO transactions (user_id, type, amount_sats, direction, status, reference, note)
		VALUES ($1, 'deposit', $2, 'in', 'pending', $3, 'Mpesa deposit — awaiting confirmation')`,
		u.ID, sats, mpesaRef)

	db.DB.Exec(`INSERT INTO audit_logs (actor_id, action, detail)
		VALUES ($1, 'deposit_requested', $2)`, u.ID, "KES "+kesStr)

	http.Redirect(w, r, "/customer/transactions?msg=Deposit+requested", http.StatusSeeOther)
}

func HandleWithdraw(w http.ResponseWriter, r *http.Request) {
	u := middleware.GetUser(r)
	var balanceSats int64
	var kesPerSat float64
	db.DB.QueryRow("SELECT balance_sats FROM wallets WHERE user_id=$1", u.ID).Scan(&balanceSats)
	db.DB.QueryRow("SELECT kes_per_sat FROM pool_settings LIMIT 1").Scan(&kesPerSat)

	if r.Method == http.MethodGet {
		serveTemplateData(w, r, "web/templates/customer/withdraw.html", map[string]interface{}{
			"User": u, "BalanceSats": balanceSats, "KesPerSat": kesPerSat,
		})
		return
	}

	satsStr := r.FormValue("sats_amount")
	sats, err := strconv.ParseInt(satsStr, 10, 64)
	if err != nil || sats <= 0 {
		serveTemplateData(w, r, "web/templates/customer/withdraw.html", map[string]interface{}{
			"User": u, "BalanceSats": balanceSats, "KesPerSat": kesPerSat, "Error": "Enter a valid amount.",
		})
		return
	}
	if sats > balanceSats {
		serveTemplateData(w, r, "web/templates/customer/withdraw.html", map[string]interface{}{
			"User": u, "BalanceSats": balanceSats, "KesPerSat": kesPerSat, "Error": "Insufficient balance.",
		})
		return
	}

	phone := r.FormValue("phone")
	db.DB.Exec(`INSERT INTO transactions (user_id, type, amount_sats, direction, status, reference, note)
		VALUES ($1, 'withdrawal', $2, 'out', 'pending', $3, 'Mpesa withdrawal — pending admin approval')`,
		u.ID, sats, phone)

	http.Redirect(w, r, "/customer/transactions?msg=Withdrawal+requested", http.StatusSeeOther)
}

func HandleSend(w http.ResponseWriter, r *http.Request) {
	u := middleware.GetUser(r)
	var balanceSats int64
	db.DB.QueryRow("SELECT balance_sats FROM wallets WHERE user_id=$1", u.ID).Scan(&balanceSats)

	var kesPerSat float64
	db.DB.QueryRow("SELECT kes_per_sat FROM pool_settings LIMIT 1").Scan(&kesPerSat)

	if r.Method == http.MethodGet {
		serveTemplateData(w, r, "web/templates/customer/send.html", map[string]interface{}{
			"User": u, "BalanceSats": balanceSats, "KesPerSat": kesPerSat,
		})
		return
	}

	satsStr := r.FormValue("sats_amount")
	sats, _ := strconv.ParseInt(satsStr, 10, 64)
	invoice := r.FormValue("invoice")

	if sats <= 0 || invoice == "" {
		serveTemplateData(w, r, "web/templates/customer/send.html", map[string]interface{}{
			"User": u, "BalanceSats": balanceSats, "KesPerSat": kesPerSat, "Error": "Enter amount and Lightning invoice.",
		})
		return
	}
	if sats > balanceSats {
		serveTemplateData(w, r, "web/templates/customer/send.html", map[string]interface{}{
			"User": u, "BalanceSats": balanceSats, "KesPerSat": kesPerSat, "Error": "Insufficient balance.",
		})
		return
	}

	db.DB.Exec(`INSERT INTO transactions (user_id, type, amount_sats, direction, status, reference, note)
		VALUES ($1, 'send', $2, 'out', 'pending', $3, 'Lightning send — pending')`,
		u.ID, sats, invoice)

	http.Redirect(w, r, "/customer/transactions?msg=Payment+queued", http.StatusSeeOther)
}

func HandleReceive(w http.ResponseWriter, r *http.Request) {
	u := middleware.GetUser(r)
	var lightningAddr string
	db.DB.QueryRow("SELECT COALESCE(lightning_addr,'') FROM wallets WHERE user_id=$1", u.ID).
		Scan(&lightningAddr)

	data := map[string]interface{}{
		"User":          u,
		"LightningAddr": lightningAddr,
	}

	if amountStr := r.URL.Query().Get("amount"); amountStr != "" {
		if sats, err := strconv.ParseInt(amountStr, 10, 64); err == nil && sats > 0 {
			data["InvoiceAmount"] = sats
			data["Invoice"] = fmt.Sprintf("lnbc%dn1p3fakexxxinvoicetoreplacewithreallndintegration", sats)
		} else {
			data["Error"] = "Enter a valid amount."
		}
	}

	serveTemplateData(w, r, "web/templates/customer/receive.html", data)
}

func HandleTransactions(w http.ResponseWriter, r *http.Request) {
	u := middleware.GetUser(r)
	msg := r.URL.Query().Get("msg")

	rows, _ := db.DB.Query(`SELECT id, type, amount_sats, direction, status,
		COALESCE(reference,''), COALESCE(note,''), created_at
		FROM transactions WHERE user_id=$1 ORDER BY created_at DESC LIMIT 50`, u.ID)
	var txs []TxRow
	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			var tx TxRow
			rows.Scan(&tx.ID, &tx.Type, &tx.AmountSats, &tx.Direction,
				&tx.Status, &tx.Reference, &tx.Note, &tx.CreatedAt)
			tx.TimeAgo = utils.TimeAgo(tx.CreatedAt)
			txs = append(txs, tx)
		}
	}

	serveTemplateData(w, r, "web/templates/customer/transactions.html", map[string]interface{}{
		"User":         u,
		"Transactions": txs,
		"Message":      msg,
	})
}

func HandleInterest(w http.ResponseWriter, r *http.Request) {
	u := middleware.GetUser(r)

	var totalInterest, thisMonth int64
	db.DB.QueryRow(`SELECT COALESCE(SUM(amount_sats),0) FROM transactions
		WHERE user_id=$1 AND type='interest'`, u.ID).Scan(&totalInterest)
	db.DB.QueryRow(`SELECT COALESCE(SUM(amount_sats),0) FROM transactions
		WHERE user_id=$1 AND type='interest' AND created_at >= date_trunc('month', NOW())`, u.ID).
		Scan(&thisMonth)

	var targetYield float64
	db.DB.QueryRow("SELECT target_yield FROM pool_settings LIMIT 1").Scan(&targetYield)

	serveTemplateData(w, r, "web/templates/customer/profit_access.html", map[string]interface{}{
		"User":          u,
		"TotalInterest": totalInterest,
		"ThisMonth":     thisMonth,
		"APY":           targetYield * 100,
	})
}

func HandleSettings(w http.ResponseWriter, r *http.Request) {
	u := middleware.GetUser(r)
	msg := ""

	if r.Method == http.MethodPost {
		fullName := r.FormValue("full_name")
		db.DB.Exec("UPDATE users SET full_name=$1 WHERE id=$2", fullName, u.ID)
		msg = "Settings saved."
	}

	serveTemplateData(w, r, "web/templates/customer/settings.html", map[string]interface{}{
		"User":    u,
		"Message": msg,
	})
}
