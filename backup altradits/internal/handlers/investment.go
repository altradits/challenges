package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/altradits/altradits/internal/db"
	"github.com/altradits/altradits/internal/middleware"
	"github.com/altradits/altradits/internal/utils"
)

type LockRow struct {
	ID           string
	AmountSats   int64
	LockYears    int
	Status       string
	ProfitSats   int64
	LockedAt     time.Time
	MaturesAt    time.Time
	DaysLeft     int
	PenaltyPct   int
	TimeAgo      string
}

func HandleInvestments(w http.ResponseWriter, r *http.Request) {
	u := middleware.GetUser(r)
	msg := r.URL.Query().Get("msg")

	rows, _ := db.DB.Query(`SELECT id, amount_sats, lock_years, status, profit_sats, locked_at, matures_at
		FROM investment_locks WHERE user_id=$1 ORDER BY locked_at DESC`, u.ID)
	var locks []LockRow
	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			var l LockRow
			rows.Scan(&l.ID, &l.AmountSats, &l.LockYears, &l.Status,
				&l.ProfitSats, &l.LockedAt, &l.MaturesAt)
			l.DaysLeft = utils.DaysUntil(l.MaturesAt)
			l.PenaltyPct = utils.EarlyWithdrawalPenalty(l.LockedAt)
			l.TimeAgo = utils.TimeAgo(l.LockedAt)
			locks = append(locks, l)
		}
	}

	var balanceSats int64
	db.DB.QueryRow("SELECT balance_sats FROM wallets WHERE user_id=$1", u.ID).Scan(&balanceSats)

	serveTemplateData(w, r, "web/templates/customer/investments.html", map[string]interface{}{
		"User":        u,
		"Locks":       locks,
		"BalanceSats": balanceSats,
		"Message":     msg,
	})
}

func HandleAddInvestment(w http.ResponseWriter, r *http.Request) {
	u := middleware.GetUser(r)

	var balanceSats int64
	db.DB.QueryRow("SELECT balance_sats FROM wallets WHERE user_id=$1", u.ID).Scan(&balanceSats)

	if r.Method == http.MethodGet {
		serveTemplateData(w, r, "web/templates/customer/add_investment.html", map[string]interface{}{
			"User": u, "BalanceSats": balanceSats,
		})
		return
	}

	satsStr := r.FormValue("sats_amount")
	yearsStr := r.FormValue("lock_years")

	sats, _ := strconv.ParseInt(satsStr, 10, 64)
	years, _ := strconv.Atoi(yearsStr)

	if sats <= 0 || sats > balanceSats {
		serveTemplateData(w, r, "web/templates/customer/add_investment.html", map[string]interface{}{
			"User": u, "BalanceSats": balanceSats, "Error": "Invalid amount or insufficient balance.",
		})
		return
	}
	if years < 5 {
		years = 5
	}

	maturesAt := time.Now().AddDate(years, 0, 0)

	// Deduct from wallet
	db.DB.Exec("UPDATE wallets SET balance_sats=balance_sats-$1, updated_at=NOW() WHERE user_id=$2", sats, u.ID)

	// Create lock
	db.DB.Exec(`INSERT INTO investment_locks (user_id, amount_sats, lock_years, matures_at)
		VALUES ($1, $2, $3, $4)`, u.ID, sats, years, maturesAt)

	// Record transaction
	db.DB.Exec(`INSERT INTO transactions (user_id, type, amount_sats, direction, status, note)
		VALUES ($1, 'investment_lock', $2, 'out', 'confirmed', $3)`,
		u.ID, sats, strconv.Itoa(years)+"-year lock created")

	http.Redirect(w, r, "/customer/investments?msg=Investment+locked+successfully", http.StatusSeeOther)
}
