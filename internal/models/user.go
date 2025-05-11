package models

import "time"

type User struct {
	Id           int       `json:"-"`
	PhoneNumber  string    `gorm:"unique;not null" json:"phone_number"`
	Username     string    `json:"username"`
	BankCard     string    `json:"bank_card"`
	Email        *string   `gorm:"unique" json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
