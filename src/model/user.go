package model

import (
    "github.com/jinzhu/gorm"
    "time"
    
)

type User struct {
	gorm.Model
	ID        int       `json:"id" db:"id"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Password  string    `json:"password" db:"password"`
	Email     string    `json:"email" db:"email"`
	Phone     string    `json:"phone" db:"phone"`
	UserType  string    `json:"user_type" db:"user_type"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
 }
 
 
 var db *gorm.DB
 