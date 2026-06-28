package models

import (
	"time"
	"github.com/google/uuid"
)

type Account struct {
	ID            uuid.UUID `json:"id"`
	OwnerName     string    `json:"owner_name"`
	Email         string    `json:"email"`
	AccountNumber string    `json:"account_number"`
	Balance       float64   `json:"balance"`
	Currency      string    `json:"currency"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Transaction struct {
	ID          uuid.UUID `json:"id"`
	AccountID   uuid.UUID `json:"account_id"`
	Type        string    `json:"type"`
	Amount      float64   `json:"amount"`
	Balance     float64   `json:"balance_after"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateAccountRequest struct {
	OwnerName string `json:"owner_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Currency  string `json:"currency" binding:"required"`
}

type DepositRequest struct {
	Amount      float64 `json:"amount" binding:"required,gt=0"`
	Description string  `json:"description"`
}

type WithdrawRequest struct {
	Amount      float64 `json:"amount" binding:"required,gt=0"`
	Description string  `json:"description"`
}
