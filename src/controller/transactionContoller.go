package controller

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"

    "github.com/gbubemi22/transaction/src/model"
    "github.com/jinzhu/gorm"
)

// ErrorResponse is a struct for handling error responses.
type ErrorResponses struct {
    Error string `json:"error"`
}

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
        }

        if err := c.ShouldBindJSON(&transferData); err != nil {
            // Override the status code to 201
            c.AbortWithStatus(http.StatusCreated)
            c.JSON(http.StatusCreated, ErrorResponses{Error: "Invalid JSON data: " + err.Error()})
            return
        }

        // Find the sender and receiver wallets
        var senderWallet model.Wallet
        var receiverWallet model.Wallet

        // Receiver
        creditReference := fmt.Sprintf("TRFR FROM: %s. TRNX REF: TRNX-%s", transferData.FromUsername, uuid.New().String())

        // Sender
        debitReference := fmt.Sprintf("TRFR TO: %s. TRNX REF: TRNX-%s", transferData.FromUsername, uuid.New().String())

        fmt.Println("Check:", transferData)

        // Check if wallet is available
        if err := tc.db.Where("user_name = ?", transferData.FromUsername).First(&senderWallet).Error; err != nil {
            tx.Rollback()
            c.JSON(http.StatusNotFound, ErrorResponses{Error: fmt.Sprintf("Sender with username %s not found", transferData.FromUsername)})
            return
        }

        // Check if wallet is available
        if err := tc.db.Where("user_name = ?", transferData.ToUsername).First(&receiverWallet).Error; err != nil {
            tx.Rollback()
            c.JSON(http.StatusNotFound, ErrorResponses{Error: fmt.Sprintf("Receiver with username %s not found", transferData.ToUsername)})
            return
        }

        // Check sender's balance
        if senderWallet.Balance < transferData.Amount {
            // Set the status code to 400 Bad Request
            c.AbortWithStatus(http.StatusBadRequest)

            // Return a JSON error response
            c.JSON(http.StatusBadRequest, ErrorResponses{Error: "Insufficient balance, controller"})
            return
        }

        // Perform the credit transaction for receiver
        receiverWallet.Balance += transferData.Amount // Update the wallet balance
        if err := tc.db.Model(&receiverWallet).Update("balance", receiverWallet.Balance).Error; err != nil {
            c.JSON(http.StatusInternalServerError, ErrorResponses{Error: "Failed to update receiver's wallet balance"})
            return
        }

        // Decrease the Sender's Wallet
        // Create transaction records
        transactionSender := model.Transaction{
            TrnxType:        "DR",
            Purpose:         "payment",
            Amount:          transferData.Amount,
            WalletUsername:  transferData.FromUsername,
            Reference:       creditReference,
            BalanceBefore:   senderWallet.Balance - transferData.Amount,
            BalanceAfter:    senderWallet.Balance,
            Summary:         transferData.Summary,
            TrnxSummary:     fmt.Sprintf("TRFR TO: %s. TRNX REF: TRNX-%s", transferData.FromUsername, uuid.New().String()),
        }
        

        // Decrease the Sender's Wallet balance
        senderWallet.Balance -= transferData.Amount
        if err := tc.db.Model(&senderWallet).Update("balance", senderWallet.Balance).Error; err != nil {
            c.JSON(http.StatusInternalServerError, ErrorResponses{Error: "Failed to update receiver's wallet balance"})
            return
        }

        // Increase the Receiver's Wallet
        transactionReceiver := model.Transaction{
            TrnxType:        "CR",
            Purpose:         "payment",
            Amount:          transferData.Amount,
            WalletUsername:  transferData.ToUsername,
            Reference:       debitReference,
            BalanceBefore:   receiverWallet.Balance - transferData.Amount,
            BalanceAfter:    receiverWallet.Balance,
            Summary:         transferData.Summary,
            TrnxSummary:     fmt.Sprintf("TRFR FROM: %s. TRNX REF: TRNX-%s", transferData.ToUsername, uuid.New().String()),
        }

        if err := tc.db.Create(&transactionSender).Error; err != nil {
            tx.Rollback()
            c.JSON(http.StatusInternalServerError, ErrorResponses{Error: "Failed to create sender's transaction record"})
            return
        }

        if err := tc.db.Create(&transactionReceiver).Error; err != nil {
            tx.Rollback()
            c.JSON(http.StatusInternalServerError, ErrorResponses{Error: "Failed to create receiver's transaction record"})
            return
        }

        // Commit the transaction
        tx.Commit()

        // Set the status code to 201 Created
        c.Status(http.StatusCreated)

        c.JSON(http.StatusCreated, gin.H{"message": "Transfer successful"})
    }
}
