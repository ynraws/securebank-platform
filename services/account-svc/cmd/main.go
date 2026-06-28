package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/securebank/account-svc/config"
	"github.com/securebank/account-svc/internal/handlers"
	"github.com/securebank/account-svc/internal/middleware"
	"github.com/securebank/account-svc/internal/repository"
)

func main() {
	// Load config
	cfg := config.Load()

	// Connect DB
	db, err := repository.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Run migrations
	repository.RunMigrations(db)

	// Setup router
	r := gin.Default()
	r.Use(middleware.RequestLogger())
	r.Use(middleware.Recovery())

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy", "service": "account-svc"})
	})

	// Routes
	h := handlers.NewAccountHandler(db)
	v1 := r.Group("/api/v1")
	{
		v1.POST("/accounts", h.CreateAccount)
		v1.GET("/accounts/:id", middleware.AuthRequired(), h.GetAccount)
		v1.GET("/accounts/:id/balance", middleware.AuthRequired(), h.GetBalance)
		v1.POST("/accounts/:id/deposit", middleware.AuthRequired(), h.Deposit)
		v1.POST("/accounts/:id/withdraw", middleware.AuthRequired(), h.Withdraw)
		v1.GET("/accounts/:id/transactions", middleware.AuthRequired(), h.GetTransactions)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Printf("account-svc starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
