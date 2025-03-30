package dto

import "time"

type Car struct {
	ID               uint        `gorm:"primaryKey" json:"id,omitempty"`
	UserID           uint        `gorm:"not null" json:"user_id,omitempty"`
	CategoryID       uint        `gorm:"not null" json:"category_id,omitempty"`
	BrandID          uint        `gorm:"not null" json:"brand_id,omitempty"`
	ColorID          uint        `gorm:"not null" json:"color_id,omitempty"`
	GenerationID     uint        `gorm:"not null" json:"generation_id,omitempty"`
	BodyID           uint        `gorm:"not null" json:"body_id,omitempty"`
	CityID           uint        `gorm:"not null" json:"city_id,omitempty"`
	ModelID          uint        `gorm:"not null" json:"model_id,omitempty"`
	User             *User       `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Category         *Category   `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Brand            *Brand      `gorm:"foreignKey:BrandID" json:"brand,omitempty"`
	Color            *Color      `gorm:"foreignKey:ColorID" json:"color,omitempty"`
	Generation       *Generation `gorm:"foreignKey:GenerationID" json:"generation,omitempty"`
	Body             *Body       `gorm:"foreignKey:BodyID" json:"body,omitempty"`
	City             *City       `gorm:"foreignKey:CityID" json:"city,omitempty"`
	Model            *Model      `gorm:"foreignKey:ModelID" json:"model,omitempty"`
	AvatarSource     string      `json:"avatar_source,omitempty"`
	EngineVolume     string      `json:"engine_volume,omitempty"`
	Mileage          string      `json:"mileage,omitempty"`
	WheelDrive       string      `json:"wheel_drive,omitempty"`
	SteeringWheel    string      `json:"steering_wheel,omitempty"`
	CustomsClearance bool        `json:"customs_clearance,omitempty"`
	Description      string      `json:"description,omitempty"`
	Price            string      `json:"price,omitempty"`
	CreatedAt        time.Time   `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt        time.Time   `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
}
