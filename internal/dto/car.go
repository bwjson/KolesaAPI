package dto

import "time"

type WheelDrive string
type SteeringWheel string

const (
	Manual    WheelDrive = "manual"
	Automatic WheelDrive = "automatic"
)

const (
	Left  SteeringWheel = "left"
	Right SteeringWheel = "right"
)

//type Car struct {
//	Id               int           `json:"id,omitempty"`
//	UserId           int           `json:"user_id,omitempty"`
//	CategoryId       int           `json:"category_id,omitempty"`
//	BrandId          int           `json:"brand_id,omitempty"`
//	ColorId          int           `json:"color_id,omitempty"`
//	GenerationId     int           `json:"generation_id,omitempty"`
//	BodyId           int           `json:"body_id,omitempty"`
//	CityId           int           `json:"city_id,omitempty"`
//	EngineVolume     string        `json:"engine_volume"`
//	Mileage          string        `json:"mileage"`
//	WheelDrive       WheelDrive    `json:"wheel_drive"`
//	SteeringWheel    SteeringWheel `json:"steering_wheel"`
//	CustomsClearance bool          `json:"customs_clearance"`
//	Description      string        `json:"description"`
//	Price            string        `json:"price"`
//	CreatedAt        time.Time     `json:"created_at,omitempty"`
//	UpdatedAt        time.Time     `json:"updated_at,omitempty"`
//}

type Car struct {
	ID               uint       `gorm:"primaryKey" json:"id"`
	UserID           uint       `json:"user_id"`
	CategoryID       uint       `json:"category_id"`
	BrandID          uint       `json:"brand_id"`
	ColorID          uint       `json:"color_id"`
	GenerationID     uint       `json:"generation_id"`
	BodyID           uint       `json:"body_id"`
	CityID           uint       `json:"city_id"`
	User             User       `gorm:"foreignKey:UserID" json:"user"`
	Category         Category   `gorm:"foreignKey:CategoryID" json:"category"`
	Brand            Brand      `gorm:"foreignKey:BrandID" json:"brand"`
	Color            Color      `gorm:"foreignKey:ColorID" json:"color"`
	Generation       Generation `gorm:"foreignKey:GenerationID" json:"generation"`
	Body             Body       `gorm:"foreignKey:BodyID" json:"body"`
	City             City       `gorm:"foreignKey:CityID" json:"city"`
	EngineVolume     string     `json:"engine_volume"`
	Mileage          string     `json:"mileage"`
	WheelDrive       string     `json:"wheel_drive"`
	SteeringWheel    string     `json:"steering_wheel"`
	CustomsClearance bool       `json:"customs_clearance"`
	Description      string     `json:"description"`
	Price            string     `json:"price"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

//type CarExtended struct {
//	Id               int           `json:"id,omitempty"`
//	User             User          `json:"user,omitempty"`
//	Category         Category      `json:"category,omitempty"`
//	Brand            Brand         `json:"brand,omitempty"`
//	Color            Color         `json:"color,omitempty"`
//	Generation       Generation    `json:"generation,omitempty"`
//	Body             Body          `json:"body,omitempty"`
//	City             City          `json:"city,omitempty"`
//	EngineVolume     string        `json:"engine_volume"`
//	Mileage          string        `json:"mileage"`
//	WheelDrive       WheelDrive    `json:"wheel_drive"`
//	SteeringWheel    SteeringWheel `json:"steering_wheel"`
//	CustomsClearance bool          `json:"customs_clearance"`
//	Description      string        `json:"description"`
//	Price            string        `json:"price"`
//	CreatedAt        time.Time     `json:"created_at,omitempty"`
//	UpdatedAt        time.Time     `json:"updated_at,omitempty"`
//}
