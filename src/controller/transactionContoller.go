package controller

import (
	"fmt"
	"net/http"
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
	"github.com/gbubemi22/transaction/src/model"
	"github.com/gbubemi22/transaction/src/helpers"
	
	"github.com/jinzhu/gorm"
)

type TransactionController struct {
	db *gorm.DB
}

func NewTransactionController(db *gorm.DB) *TransactionController {
	return &TransactionController{db}
}

func (tc *TransactionController) Transfer() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start a new transaction
		tx := tc.db.Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		var transferData struct {
			ToUsername   string  `json:"toUsername"`
			FromUsername string  `json:"fromUsername"`
			Amount       float64 `json:"amount"`
			Summary      string  `json:"summary"`
			TranxType    string  `json:"tranxType"`
		}

		if err := c.ShouldBindJSON(&transferData); err != nil {
			c.JSON(http.StatusBadRequest, helpers.ErrorResponse{Error: "Invalid JSON data: " + err.Error()})
			return
		}

		// Find the sender and receiver wallets
		var senderWallet model.Wallet
		var receiverWallet model.Wallet

		if err := tc.db.Where("username = ?", transferData.FromUsername).First(&senderWallet).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusNotFound, helpers.ErrorResponse{Error: fmt.Sprintf("Sender with username %s not found", transferData.FromUsername)})
			return
		}

		if err := tc.db.Where("username = ?", transferData.ToUsername).First(&receiverWallet).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusNotFound, helpers.ErrorResponse{Error: fmt.Sprintf("Receiver with username %s not found", transferData.ToUsername)})
			return
		}

		// Check sender's balance
		if senderWallet.Balance < transferData.Amount {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, helpers.ErrorResponse{Error: "Insufficient balance"})
			return
		}

		// Generate UUID references
		creditReference := fmt.Sprintf("TRFR TO: %s. TRNX REF: TRNX-%s", transferData.ToUsername, uuid.New().String())
		 debitReference := fmt.Sprintf("TRFR FROM: %s. TRNX REF: TRNX-%s", transferData.FromUsername, uuid.New().String())

		// Deduct from sender and credit to receiver using the credit and debit helper functions
		helpers.CreditAccount(c, tc.db)
		helpers.DebitAccount(c, tc.db)


		// Create transaction records
		transactionSender := model.Transaction{
			TrnxType:       transferData.TranxType,
			Purpose:        "payment",
			Amount:         transferData.Amount,
			WalletUsername:       transferData.FromUsername,
			Reference:      creditReference,
			BalanceBefore:  senderWallet.Balance + transferData.Amount,
			BalanceAfter:   senderWallet.Balance,
			Summary:        transferData.Summary,
			TrnxSummary:    fmt.Sprintf("TRFR TO: %s. TRNX REF: TRNX-%s", transferData.ToUsername, uuid.New().String()), 
		}

		transactionReceiver := model.Transaction{
			TrnxType:       transferData.TranxType,
			Purpose:        "payment",
			Amount:         transferData.Amount,
			WalletUsername:       transferData.ToUsername,
			Reference:  debitReference,
			BalanceBefore:  receiverWallet.Balance - transferData.Amount,
			BalanceAfter:   receiverWallet.Balance,
			Summary:        transferData.Summary,
			TrnxSummary:    fmt.Sprintf("TRFR FROM: %s. TRNX REF: TRNX-%s", transferData.FromUsername, uuid.New().String()), 
		}

		if err := tc.db.Create(&transactionSender).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{Error: "Failed to create sender's transaction record"})
			return
		}

		if err := tc.db.Create(&transactionReceiver).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{Error: "Failed to create receiver's transaction record"})
			return
		}

		// Commit the transaction
		tx.Commit()

		c.JSON(http.StatusCreated, gin.H{"message": "Transfer successful"})
	}
}
