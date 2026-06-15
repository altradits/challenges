package handlers

import (
	"net/http"

	"github.com/altradits/altradits/internal/db"
	"github.com/altradits/altradits/internal/middleware"
)

type AssetRow struct {
	ID            string
	Ticker        string
	Name          string
	Category      string
	AllocationPct float64
	CurrentValue  int64
	YieldPct      float64
}

func HandleTraderDashboard(w http.ResponseWriter, r *http.Request) {
	u := middleware.GetUser(r)

	var totalAUM int64
	db.DB.QueryRow("SELECT COALESCE(SUM(balance_sats),0) FROM wallets").Scan(&totalAUM)

	var targetYield, adminFee float64
	db.DB.QueryRow("SELECT target_yield, admin_fee_pct FROM pool_settings LIMIT 1").
		Scan(&targetYield, &adminFee)

	rows, _ := db.DB.Query(`SELECT id, ticker, name, category, allocation_pct, current_value, yield_pct
		FROM assets ORDER BY allocation_pct DESC`)
	var assets []AssetRow
	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			var a AssetRow
			rows.Scan(&a.ID, &a.Ticker, &a.Name, &a.Category,
				&a.AllocationPct, &a.CurrentValue, &a.YieldPct)
			assets = append(assets, a)
		}
	}

	var totalInterestPaid int64
	db.DB.QueryRow("SELECT COALESCE(SUM(amount_sats),0) FROM transactions WHERE type='interest'").Scan(&totalInterestPaid)

	serveTemplateData(w, r, "web/templates/trader/dashboard.html", map[string]interface{}{
		"User":             u,
		"TotalAUM":         totalAUM,
		"TargetYield":      targetYield * 100,
		"AdminFee":         adminFee * 100,
		"Assets":           assets,
		"TotalInterestPaid": totalInterestPaid,
	})
}

func HandleAddAsset(w http.ResponseWriter, r *http.Request) {
	u := middleware.GetUser(r)

	if r.Method == http.MethodGet {
		serveTemplateData(w, r, "web/templates/trader/add_asset.html", map[string]interface{}{"User": u})
		return
	}

	ticker := r.FormValue("ticker")
	name := r.FormValue("name")
	category := r.FormValue("category")
	allocation := r.FormValue("allocation_pct")
	yield := r.FormValue("yield_pct")

	db.DB.Exec(`INSERT INTO assets (ticker, name, category, allocation_pct, yield_pct)
		VALUES ($1, $2, $3, $4, $5)`, ticker, name, category, allocation, yield)

	http.Redirect(w, r, "/trader/dashboard", http.StatusSeeOther)
}

func HandleProfitUpdate(w http.ResponseWriter, r *http.Request) {
	u := middleware.GetUser(r)

	if r.Method == http.MethodGet {
		serveTemplateData(w, r, "web/templates/trader/profit_update.html", map[string]interface{}{"User": u})
		return
	}

	sats := r.FormValue("amount_sats")
	note := r.FormValue("note")

	db.DB.Exec(`INSERT INTO profit_logs (amount_sats, note, created_by) VALUES ($1, $2, $3)`,
		sats, note, u.ID)

	http.Redirect(w, r, "/trader/dashboard?msg=Profit+recorded", http.StatusSeeOther)
}
