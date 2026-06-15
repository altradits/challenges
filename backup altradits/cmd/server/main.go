package main

import (
	"log"
	"net/http"
	"os"

	"github.com/altradits/altradits/internal/db"
	"github.com/altradits/altradits/internal/handlers"
	"github.com/altradits/altradits/internal/middleware"
)

func main() {
	// Connect to database
	if err := db.Connect(); err != nil {
		log.Fatalf("[FATAL] database: %v", err)
	}

	// Run migrations
	if err := db.RunMigrations(); err != nil {
		log.Fatalf("[FATAL] migrations: %v", err)
	}

	// Seed default data
	if err := db.Seed(); err != nil {
		log.Printf("[WARN] seed: %v", err)
	}

	mux := http.NewServeMux()

	// ── Static assets ──────────────────────────────────────────────────────
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// ── Public routes ───────────────────────────────────────────────────────
	mux.HandleFunc("/", handlers.HandleHome)
	mux.HandleFunc("/register", handlers.HandleRegister)
	mux.HandleFunc("/login", handlers.HandleLogin)
	mux.HandleFunc("/logout", handlers.HandleLogout)

	// ── Customer routes ─────────────────────────────────────────────────────
	mux.HandleFunc("/customer/dashboard", middleware.RequireAuth(handlers.HandleCustomerDashboard))
	mux.HandleFunc("/customer/deposit", middleware.RequireAuth(handlers.HandleDeposit))
	mux.HandleFunc("/customer/withdraw", middleware.RequireAuth(handlers.HandleWithdraw))
	mux.HandleFunc("/customer/send", middleware.RequireAuth(handlers.HandleSend))
	mux.HandleFunc("/customer/receive", middleware.RequireAuth(handlers.HandleReceive))
	mux.HandleFunc("/customer/transactions", middleware.RequireAuth(handlers.HandleTransactions))
	mux.HandleFunc("/customer/interest", middleware.RequireAuth(handlers.HandleInterest))
	mux.HandleFunc("/customer/investments", middleware.RequireAuth(handlers.HandleInvestments))
	mux.HandleFunc("/customer/investments/new", middleware.RequireAuth(handlers.HandleAddInvestment))
	mux.HandleFunc("/customer/settings", middleware.RequireAuth(handlers.HandleSettings))

	// ── Trader routes ────────────────────────────────────────────────────────
	mux.HandleFunc("/trader/dashboard", middleware.RequireRole("trader", handlers.HandleTraderDashboard))
	mux.HandleFunc("/trader/assets/new", middleware.RequireRole("trader", handlers.HandleAddAsset))
	mux.HandleFunc("/trader/profit", middleware.RequireRole("trader", handlers.HandleProfitUpdate))

	// ── Agent routes ─────────────────────────────────────────────────────────
	mux.HandleFunc("/agent/dashboard", middleware.RequireRole("agent", handlers.HandleAgentDashboard))
	mux.HandleFunc("/agent/load", middleware.RequireRole("agent", handlers.HandleAgentCashIn))
	mux.HandleFunc("/agent/cashout", middleware.RequireRole("agent", handlers.HandleAgentCashOut))
	mux.HandleFunc("/agent/transactions", middleware.RequireRole("agent", handlers.HandleAgentTransactions))

	// ── Admin routes ─────────────────────────────────────────────────────────
	mux.HandleFunc("/admin/dashboard", middleware.RequireRole("admin", handlers.HandleAdminDashboard))
	mux.HandleFunc("/admin/customers", middleware.RequireRole("admin", handlers.HandleAdminCustomers))
	mux.HandleFunc("/admin/deposits", middleware.RequireRole("admin", handlers.HandleAdminDeposits))
	mux.HandleFunc("/admin/withdrawals", middleware.RequireRole("admin", handlers.HandleAdminWithdrawals))
	mux.HandleFunc("/admin/settings", middleware.RequireRole("admin", handlers.HandleAdminSettings))
	mux.HandleFunc("/admin/users/block", middleware.RequireRole("admin", handlers.HandleAdminBlockUser))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	log.Printf("  Altradits running on http://localhost:%s", port)
	log.Printf("  Admin: admin@altradits.com / admin123")
	log.Printf("  Trader: trader@altradits.com / trader123")
	log.Printf("  Agent: agent@altradits.com / agent123")
	log.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("server: %v", err)
	}
}
