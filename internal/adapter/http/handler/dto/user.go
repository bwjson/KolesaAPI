package dto

import "time"

type SafeUserDTO struct {
	PhoneNumber string    `gorm:"unique;not null" json:"phone_number"`
	Username    string    `json:"username"`
	Email       *string   `gorm:"unique" json:"email"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
