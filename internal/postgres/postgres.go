package postgres

import (
	"fmt"
	"github.com/bwjson/kolesa_api/internal/dto"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func NewPostgresDB(user string, name string, port string, password string, host string, sslmode string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("port=%s user=%s password=%s host=%s dbname=%s sslmode=%s", port, user, password, host, name, sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&dto.User{},
		&dto.Category{},
		&dto.Brand{},
		&dto.Color{},
		&dto.Generation{},
		&dto.Body{},
		&dto.City{},
		&dto.Car{},
		&dto.CarPhoto{},
		&dto.Model{},
	)
	if err != nil {
		log.Fatal("Migration failed", err)
		return nil, err
	}

	return db, nil
}
