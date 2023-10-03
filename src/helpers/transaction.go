package helpers

import (
    "github.com/gin-gonic/gin"
  
    "github.com/jinzhu/gorm"
    "github.com/gbubemi22/transaction/src/model"
    "net/http"
    "fmt"
)

// ErrorResponse is a struct for handling error responses.
type ErrorResponse struct {
    Error string `json:"error"`
}

func CreditAccount(c *gin.Context, db *gorm.DB) {
	var creditData struct {
	    Amount     float64 `json:"amount" binding:"required"`
	    Username   string  `json:"username" binding:"required"`
	    Purpose    string  `json:"purpose" binding:"required"`
	    Reference  string  `json:"reference"`
	    Summary    string  `json:"summary"`
	    TrnxSummary string `json:"trnxSummary"`
	}
 
	// Parse credit transaction data from the request body
	if err := c.ShouldBindJSON(&creditData); err != nil {
	    c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid JSON data: " + err.Error()})
	    return
	}
 
	// Find the user's wallet by username
	var wallet model.Wallet
	if err := db.Where("username = ?", creditData.Username).First(&wallet).Error; err != nil {
	    c.JSON(http.StatusNotFound, ErrorResponse{Error: fmt.Sprintf("User %s doesn't exist", creditData.Username)})
	    return
	}
 
	// Perform the credit transaction
	wallet.Balance += creditData.Amount
	if err := db.Save(&wallet).Error; err != nil {
	    c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to update wallet balance"})
	    return
	}
 
	// Create a transaction record
	transaction := model.Transaction{
	    TrnxType:       "CR",
	    Purpose:        creditData.Purpose,
	    Amount:         creditData.Amount,
	    Reference:      creditData.Reference,
	    BalanceBefore:  wallet.Balance - creditData.Amount,
	    BalanceAfter:   wallet.Balance,
	    Summary:        creditData.Summary,
	    TrnxSummary:    creditData.TrnxSummary,
	}
 
	if err := db.Create(&transaction).Error; err != nil {
	    c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to create transaction"})
	    return
	}
 
	c.JSON(http.StatusOK, gin.H{"message": "Credit successful", "wallet": wallet, "transaction": transaction})
}

func DebitAccount(c *gin.Context, db *gorm.DB) {
	var debitData struct {
	    Amount     float64 `json:"amount"`
	    Username   string  `json:"username"`
	    Purpose    string  `json:"purpose"`
	    Reference  string  `json:"reference"`
	    Summary    string  `json:"summary"`
	    TrnxSummary string `json:"trnxSummary"`
	}

	// Parse debit transaction data from the request body
	if err := c.ShouldBindJSON(&debitData); err != nil {
	    c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid JSON data: " + err.Error()})
	    return
	}

	// Find the user's wallet by username
	var wallet model.Wallet
	if err := db.Where("username = ?", debitData.Username).First(&wallet).Error; err != nil {
	    c.JSON(http.StatusNotFound, ErrorResponse{Error: fmt.Sprintf("User %s doesn't exist", debitData.Username)})
	    return
	}

	// Check if the user has sufficient balance
	if wallet.Balance < debitData.Amount {
	    c.JSON(http.StatusBadRequest, ErrorResponse{Error: fmt.Sprintf("User %s has insufficient balance", debitData.Username)})
	    return
	}

	// Perform the debit transaction
	wallet.Balance -= debitData.Amount
	if err := db.Save(&wallet).Error; err != nil {
	    c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to update wallet balance"})
	    return
	}

	// Create a transaction record
	transaction := model.Transaction{
	    TrnxType:       "DR",
	    Purpose:        debitData.Purpose,
	    Amount:         debitData.Amount,
	    Reference:      debitData.Reference,
	    BalanceBefore:  wallet.Balance + debitData.Amount,
	    BalanceAfter:   wallet.Balance,
	    Summary:        debitData.Summary,
	    TrnxSummary:    debitData.TrnxSummary,
	}

	if err := db.Create(&transaction).Error; err != nil {
	    c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to create transaction"})
	    return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Debit successful", "wallet": wallet, "transaction": transaction})
}
