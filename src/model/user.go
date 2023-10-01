package model

import (
    "github.com/jinzhu/gorm"
    
)

type User struct {
    gorm.Model
    FirstName string    `json:"first_name" db:"first_name"`
    LastName  string    `json:"last_name" db:"last_name"`
    Password  string    `json:"password" db:"password"`
    Email     string    `json:"email" db:"email"`
    Phone     string    `json:"phone" db:"phone"`
    UserType  string    `json:"user_type" db:"user_type"`
   
}

func (u *User) CreateUser(db *gorm.DB) error {
    return db.Create(u).Error
}

func (u *User) GetUserByID(db *gorm.DB, id uint) error {
    return db.First(u, id).Error
}


func (u *User) GetUserByEmail(db *gorm.DB, email string) error {
    return db.Where("email = ?", email).First(u).Error
}

func (u *User) GetUserByPhone(db *gorm.DB, phone string) error {
    return db.Where("phone = ?", phone).First(u).Error
}

func (u *User) UpdateUser(db *gorm.DB) error {
    return db.Save(u).Error
}

func (u *User) DeleteUser(db *gorm.DB) error {
    return db.Delete(u).Error
}

func GetAllUsers(db *gorm.DB) ([]User, error) {
    var users []User
    err := db.Find(&users).Error
    if err != nil {
        return nil, err
    }
    return users, nil
}
