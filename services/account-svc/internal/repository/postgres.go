package repository

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

func NewPostgresDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	log.Println("Connected to PostgreSQL successfully")
	return db, nil
}

func RunMigrations(db *sql.DB) {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS accounts (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			owner_name VARCHAR(255) NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			account_number VARCHAR(20) UNIQUE NOT NULL,
			balance DECIMAL(15,2) DEFAULT 0.00,
			currency VARCHAR(3) DEFAULT 'INR',
			status VARCHAR(20) DEFAULT 'active',
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW()
		)`,
		`CREATE TABLE IF NOT EXISTS transactions (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			account_id UUID REFERENCES accounts(id),
			type VARCHAR(20) NOT NULL,
			amount DECIMAL(15,2) NOT NULL,
			balance_after DECIMAL(15,2) NOT NULL,
			description TEXT,
			created_at TIMESTAMP DEFAULT NOW()
		)`,
		`CREATE INDEX IF NOT EXISTS idx_transactions_account_id ON transactions(account_id)`,
	}
	for _, q := range queries {
		if _, err := db.Exec(q); err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
	}
	log.Println("Migrations completed successfully")
}
