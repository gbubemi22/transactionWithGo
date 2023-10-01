package model

import (
    "github.com/jinzhu/gorm"
)

type Wallet struct {
    gorm.Model
    UserName string  `json:"username" db:"username" required:"true"`
    Balance  float64 `json:"balance" db:"balance" required:"true" default:"0.00"`
}


func (w *Wallet) CreateUser(db *gorm.DB) error {
	return db.Create(w).Error
 }
 

 func GetWalletByID(db *gorm.DB, id uint) (*Wallet, error) {
	var wallet Wallet
	if err := db.Where("id = ?", id).First(&wallet).Error; err != nil {
	    return nil, err 
	}
	return &wallet, nil
 }
 
 // GetWalletByUserName retrieves a wallet by its UserName.
 func GetWalletByUserName(db *gorm.DB, userName string) (*Wallet, error) {
	var wallet Wallet
	if err := db.Where("username = ?", userName).First(&wallet).Error; err != nil {
	    return nil, err 
	}
	return &wallet, nil
 }
 
  