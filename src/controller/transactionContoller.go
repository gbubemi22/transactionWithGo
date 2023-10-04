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
            c.AbortWithStatus(http.StatusBadRequest)
            c.JSON(http.StatusBadRequest, ErrorResponses{Error: "Insufficient balance"})
            return
        }

        // Credit receiver's account
        CreditAccount(c, tc.db, transferData.ToUsername, transferData.Amount, creditReference, transferData.Summary)

        // Debit sender's account
        DebitAccount(c, tc.db, transferData.FromUsername, transferData.Amount, debitReference, transferData.Summary)

        // Commit the transaction
        tx.Commit()

        // Set the status code to 201 Created
        c.Status(http.StatusCreated)
        c.JSON(http.StatusCreated, gin.H{"message": "Transfer successful"})
    }
}

func CreditAccount(c *gin.Context, db *gorm.DB, username string, amount float64, reference string, summary string) {
    // Find the user's wallet by username
    var wallet model.Wallet
    if err := db.Where("user_name = ?", username).First(&wallet).Error; err != nil {
        c.JSON(http.StatusNotFound, ErrorResponses{Error: fmt.Sprintf("User %s doesn't exist", username)})
        return
    }

    // Perform the credit transaction
    wallet.Balance += amount
    if err := db.Model(&wallet).Update("balance", wallet.Balance).Error; err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponses{Error: "Failed to update wallet balance"})
        return
    }

    transaction := model.Transaction{
        TrnxType:       "CR",
        Purpose:        "transfer",
        Amount:         amount,
        Reference:      reference,
        BalanceBefore:  wallet.Balance - amount,
        BalanceAfter:   wallet.Balance,
        Summary:        summary,
        TrnxSummary:    fmt.Sprintf("TRFR TO: %s. TRNX REF: TRNX-%s", username, uuid.New().String()),
        WalletUsername: username,
    }

    if err := db.Create(&transaction).Error; err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponses{Error: "Failed to create transaction"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Credit successful", "wallet": wallet, "transaction": transaction})
}

func DebitAccount(c *gin.Context, db *gorm.DB, username string, amount float64, reference string, summary string) {
    // Find the user's wallet by username
    var wallet model.Wallet
    if err := db.Where("user_name = ?", username).First(&wallet).Error; err != nil {
        c.JSON(http.StatusNotFound, ErrorResponses{Error: fmt.Sprintf("User %s doesn't exist", username)})
        return
    }

    // Check if the user has sufficient balance
    if wallet.Balance < amount {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: fmt.Sprintf("User %s has insufficient balance", username)})
        return
    }

    // Perform the debit transaction
    wallet.Balance -= amount
    if err := db.Model(&wallet).Update("balance", wallet.Balance).Error; err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to update wallet balance"})
        return
    }

    transaction := model.Transaction{
        TrnxType:       "DR",
        Purpose:        "transfer",
        Amount:         amount,
        Reference:      reference,
        BalanceBefore:  wallet.Balance + amount,
        BalanceAfter:   wallet.Balance,
        Summary:        summary,
        TrnxSummary:    fmt.Sprintf("TRFR TO: %s. TRNX REF: TRNX-%s", username, uuid.New().String()),
        WalletUsername: username,
    }

    if err := db.Create(&transaction).Error; err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to create transaction"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Debit successful", "wallet": wallet, "transaction": transaction})
}








// package controller

// import (
//     "fmt"
//     "net/http"
//     "github.com/gin-gonic/gin"
//     "github.com/google/uuid"

//     "github.com/gbubemi22/transaction/src/model"
//     "github.com/jinzhu/gorm"
// )

// // ErrorResponse is a struct for handling error responses.
// type ErrorResponses struct {
//     Error string `json:"error"`
// }

// type TransactionController struct {
//     db *gorm.DB
// }

// func NewTransactionController(db *gorm.DB) *TransactionController {
//     return &TransactionController{db}
// }

// func (tc *TransactionController) Transfer() gin.HandlerFunc {
//     return func(c *gin.Context) {
//         // Start a new transaction
//         tx := tc.db.Begin()

//         defer func() {
//             if r := recover(); r != nil {
//                 tx.Rollback()
//             }
//         }()

//         var transferData struct {
//             ToUsername   string  `json:"toUsername"`
//             FromUsername string  `json:"fromUsername"`
//             Amount       float64 `json:"amount"`
//             Summary      string  `json:"summary"`
//         }

//         if err := c.ShouldBindJSON(&transferData); err != nil {
//             // Override the status code to 201
//             c.AbortWithStatus(http.StatusCreated)
//             c.JSON(http.StatusCreated, ErrorResponses{Error: "Invalid JSON data: " + err.Error()})
//             return
//         }

//         // Find the sender and receiver wallets
//         var senderWallet model.Wallet
//         var receiverWallet model.Wallet

//         // Receiver
//         creditReference := fmt.Sprintf("TRFR FROM: %s. TRNX REF: TRNX-%s", transferData.FromUsername, uuid.New().String())

//         // Sender
//         debitReference := fmt.Sprintf("TRFR TO: %s. TRNX REF: TRNX-%s", transferData.FromUsername, uuid.New().String())

//         fmt.Println("Check:", transferData)

//         // Check if wallet is available
//         if err := tc.db.Where("user_name = ?", transferData.FromUsername).First(&senderWallet).Error; err != nil {
//             tx.Rollback()
//             c.JSON(http.StatusNotFound, ErrorResponses{Error: fmt.Sprintf("Sender with username %s not found", transferData.FromUsername)})
//             return
//         }

//         // Check if wallet is available
//         if err := tc.db.Where("user_name = ?", transferData.ToUsername).First(&receiverWallet).Error; err != nil {
//             tx.Rollback()
//             c.JSON(http.StatusNotFound, ErrorResponses{Error: fmt.Sprintf("Receiver with username %s not found", transferData.ToUsername)})
//             return
//         }

//         // Check sender's balance
//         if senderWallet.Balance < transferData.Amount {
//             // Set the status code to 400 Bad Request
//             c.AbortWithStatus(http.StatusBadRequest)

//             // Return a JSON error response
//             c.JSON(http.StatusBadRequest, ErrorResponses{Error: "Insufficient balance, controller"})
//             return
//         }
           

//         // fix this too in the func 
//         // Perform the credit transaction for receiver
//         receiverWallet.Balance += transferData.Amount // Update the wallet balance
//         if err := tc.db.Model(&receiverWallet).Update("balance", receiverWallet.Balance).Error; err != nil {
//             c.JSON(http.StatusInternalServerError, ErrorResponses{Error: "Failed to update receiver's wallet balance"})
//             return
//         }

//         // Decrease the Sender's Wallet
//         // Create transaction records
//         transactionSender := model.Transaction{
//             TrnxType:        "DR",
//             Purpose:         "payment",
//             Amount:          transferData.Amount,
//             WalletUsername:  transferData.FromUsername,
//             Reference:       creditReference,
//             BalanceBefore:   senderWallet.Balance - transferData.Amount,
//             BalanceAfter:    senderWallet.Balance,
//             Summary:         transferData.Summary,
//             TrnxSummary:     fmt.Sprintf("TRFR TO: %s. TRNX REF: TRNX-%s", transferData.FromUsername, uuid.New().String()),
//         }
        
//           //fix this in  the func
//         // Decrease the Sender's Wallet balance
//         senderWallet.Balance -= transferData.Amount
//         if err := tc.db.Model(&senderWallet).Update("balance", senderWallet.Balance).Error; err != nil {
//             c.JSON(http.StatusInternalServerError, ErrorResponses{Error: "Failed to update receiver's wallet balance"})
//             return
//         }

//         // Increase the Receiver's Wallet
//         transactionReceiver := model.Transaction{
//             TrnxType:        "CR",
//             Purpose:         "payment",
//             Amount:          transferData.Amount,
//             WalletUsername:  transferData.ToUsername,
//             Reference:       debitReference,
//             BalanceBefore:   receiverWallet.Balance - transferData.Amount,
//             BalanceAfter:    receiverWallet.Balance,
//             Summary:         transferData.Summary,
//             TrnxSummary:     fmt.Sprintf("TRFR FROM: %s. TRNX REF: TRNX-%s", transferData.ToUsername, uuid.New().String()),
//         }

//         if err := tc.db.Create(&transactionSender).Error; err != nil {
//             tx.Rollback()
//             c.JSON(http.StatusInternalServerError, ErrorResponses{Error: "Failed to create sender's transaction record"})
//             return
//         }

//         if err := tc.db.Create(&transactionReceiver).Error; err != nil {
//             tx.Rollback()
//             c.JSON(http.StatusInternalServerError, ErrorResponses{Error: "Failed to create receiver's transaction record"})
//             return
//         }

//         // Commit the transaction
//         tx.Commit()

//         // Set the status code to 201 Created
//         c.Status(http.StatusCreated)

//         c.JSON(http.StatusCreated, gin.H{"message": "Transfer successful"})
//     }
// }
