package model 


import (
	"github.com/jinzhu/gorm"
	"github.com/google/uuid"
 )


type Transaction struct {
	gorm.Model
	TrnxType      string  `json:"trnxType" gorm:"column:trnxType"`
	Purpose       string  `json:"purpose"`
	Amount        float64 `json:"amount" gorm:"type:numeric"`
	WalletUsername string `json:"walletUsername" gorm:"column:walletUsername"`
	Reference     string  `json:"reference"`
	BalanceBefore float64 `json:"balanceBefore" gorm:"type:numeric"`
	BalanceAfter  float64 `json:"balanceAfter" gorm:"type:numeric"`
	Summary       string  `json:"summary"`
	TrnxSummary   string  `json:"trnxSummary"`
	Status        string  `json:"status" gorm:"type:varchar(255)"`
 }
 
 // GenerateUUID generates a new UUID and sets it as the reference
 func (t *Transaction) GenerateUUID() {
	uuidValue := uuid.New()
	t.Reference = uuidValue.String()
 }
 
 func (t *Transaction) InitTransaction(db *gorm.DB) error {
	t.GenerateUUID() // Generate UUID before creating the transaction
	return db.Create(t).Error
 }
 