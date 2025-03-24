package dto

import "time"

type User struct {
	Id           int       `json:"id"`
	PhoneNumber  string    `gorm:"unique;not null" json:"phone_number"`
	Username     string    `json:"username"`
	BankCard     string    `json:"bank_card"`
	Email        string    `gorm:"unique;not null" json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
