package model 


import (
	"github.com/jinzhu/gorm"
	//"github.com/google/uuid"
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
 
 
 
 func (t *Transaction) InitTransaction(db *gorm.DB) error {
	
	return db.Create(t).Error
 }
 