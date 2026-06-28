package handlers

import (
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/securebank/account-svc/internal/models"
)

type AccountHandler struct {
	db *sql.DB
}

func NewAccountHandler(db *sql.DB) *AccountHandler {
	return &AccountHandler{db: db}
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var req models.CreateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accountNumber := fmt.Sprintf("SB%d", rand.New(rand.NewSource(time.Now().UnixNano())).Intn(9000000000)+1000000000)
	var account models.Account
	err := h.db.QueryRow(`
		INSERT INTO accounts (owner_name, email, account_number, currency)
		VALUES ($1, $2, $3, $4)
		RETURNING id, owner_name, email, account_number, balance, currency, status, created_at`,
		req.OwnerName, req.Email, accountNumber, req.Currency,
	).Scan(&account.ID, &account.OwnerName, &account.Email,
		&account.AccountNumber, &account.Balance, &account.Currency,
		&account.Status, &account.CreatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create account"})
		return
	}
	c.JSON(http.StatusCreated, account)
}

func (h *AccountHandler) GetAccount(c *gin.Context) {
	id := c.Param("id")
	var account models.Account
	err := h.db.QueryRow(`
		SELECT id, owner_name, email, account_number, balance, currency, status, created_at
		FROM accounts WHERE id = $1`, id,
	).Scan(&account.ID, &account.OwnerName, &account.Email,
		&account.AccountNumber, &account.Balance, &account.Currency,
		&account.Status, &account.CreatedAt)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get account"})
		return
	}
	c.JSON(http.StatusOK, account)
}

func (h *AccountHandler) GetBalance(c *gin.Context) {
	id := c.Param("id")
	var balance float64
	err := h.db.QueryRow(`SELECT balance FROM accounts WHERE id = $1`, id).Scan(&balance)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"account_id": id, "balance": balance})
}

func (h *AccountHandler) Deposit(c *gin.Context) {
	id := c.Param("id")
	var req models.DepositRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var newBalance float64
	err := h.db.QueryRow(`
		UPDATE accounts SET balance = balance + $1, updated_at = NOW()
		WHERE id = $2 AND status = 'active'
		RETURNING balance`, req.Amount, id,
	).Scan(&newBalance)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "account not found or inactive"})
		return
	}
	h.db.Exec(`
		INSERT INTO transactions (account_id, type, amount, balance_after, description)
		VALUES ($1, 'deposit', $2, $3, $4)`,
		id, req.Amount, newBalance, req.Description)
	c.JSON(http.StatusOK, gin.H{"message": "deposit successful", "new_balance": newBalance})
}

func (h *AccountHandler) Withdraw(c *gin.Context) {
	id := c.Param("id")
	var req models.WithdrawRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var currentBalance float64
	h.db.QueryRow(`SELECT balance FROM accounts WHERE id = $1`, id).Scan(&currentBalance)
	if currentBalance < req.Amount {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insufficient balance"})
		return
	}
	var newBalance float64
	h.db.QueryRow(`
		UPDATE accounts SET balance = balance - $1, updated_at = NOW()
		WHERE id = $2 RETURNING balance`, req.Amount, id,
	).Scan(&newBalance)
	h.db.Exec(`
		INSERT INTO transactions (account_id, type, amount, balance_after, description)
		VALUES ($1, 'withdrawal', $2, $3, $4)`,
		id, req.Amount, newBalance, req.Description)
	c.JSON(http.StatusOK, gin.H{"message": "withdrawal successful", "new_balance": newBalance})
}

func (h *AccountHandler) GetTransactions(c *gin.Context) {
	id := c.Param("id")
	rows, err := h.db.Query(`
		SELECT id, account_id, type, amount, balance_after, description, created_at
		FROM transactions WHERE account_id = $1
		ORDER BY created_at DESC LIMIT 50`, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get transactions"})
		return
	}
	defer rows.Close()
	var txns []models.Transaction
	for rows.Next() {
		var t models.Transaction
		rows.Scan(&t.ID, &t.AccountID, &t.Type, &t.Amount, &t.Balance, &t.Description, &t.CreatedAt)
		txns = append(txns, t)
	}
	if txns == nil {
		txns = []models.Transaction{}
	}
	c.JSON(http.StatusOK, gin.H{"account_id": id, "transactions": txns, "count": len(txns)})
}

func generateID() string {
	return uuid.New().String()
}
