package dto

import "time"

type WheelDrive string
type SteeringWheel string

const (
	Manual    WheelDrive = "Механике"
	Automatic WheelDrive = "Автомат"
)

const (
	Left  SteeringWheel = "Левое"
	Right SteeringWheel = "Правое"
)

type Car struct {
	ID               uint          `gorm:"primaryKey" json:"id"`
	UserID           uint          `gorm:"not null" json:"user_id"`
	CategoryID       uint          `gorm:"not null" json:"category_id"`
	BrandID          uint          `gorm:"not null" json:"brand_id"`
	ColorID          uint          `gorm:"not null" json:"color_id"`
	GenerationID     uint          `gorm:"not null" json:"generation_id"`
	BodyID           uint          `gorm:"not null" json:"body_id"`
	CityID           uint          `gorm:"not null" json:"city_id"`
	User             User          `gorm:"foreignKey:UserID" json:"user"`
	Category         Category      `gorm:"foreignKey:CategoryID" json:"category"`
	Brand            Brand         `gorm:"foreignKey:BrandID" json:"brand"`
	Color            Color         `gorm:"foreignKey:ColorID" json:"color"`
	Generation       Generation    `gorm:"foreignKey:GenerationID" json:"generation"`
	Body             Body          `gorm:"foreignKey:BodyID" json:"body"`
	City             City          `gorm:"foreignKey:CityID" json:"city"`
	EngineVolume     string        `json:"engine_volume"`
	Mileage          string        `json:"mileage"`
	WheelDrive       WheelDrive    `json:"wheel_drive"`
	SteeringWheel    SteeringWheel `json:"steering_wheel"`
	CustomsClearance bool          `json:"customs_clearance"`
	Description      string        `json:"description"`
	Price            string        `json:"price"`
	CreatedAt        time.Time     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
}
