package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(user string, name string, port string, password string, host string, sslmode string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("port=%s user=%s password=%s host=%s dbname=%s sslmode=%s", port, user, password, host, name, sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
