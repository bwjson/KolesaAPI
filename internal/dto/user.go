package dto

import "time"

//CREATE TABLE users (
//id SERIAL PRIMARY KEY,
//phone_number VARCHAR(20) UNIQUE NOT NULL,
//username VARCHAR(100),
//bank_card VARCHAR(50),
//email VARCHAR(255) UNIQUE,
//password_hash TEXT NOT NULL,
//created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
//);

type User struct {
	Id           int       `json:"id"`
	PhoneNumber  string    `json:"phone_number"`
	Username     string    `json:"username"`
	BankCard     string    `json:"bank_card"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
