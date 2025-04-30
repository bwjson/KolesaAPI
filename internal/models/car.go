package models

import "time"

type Car struct {
	// Relationships IDs
	ID           uint `gorm:"primaryKey" json:"id"`
	UserID       uint `gorm:"not null" json:"-"`
	CategoryID   uint `gorm:"not null" json:"-"`
	BrandID      uint `gorm:"not null" json:"-"`
	ColorID      uint `gorm:"not null" json:"-"`
	GenerationID uint `gorm:"not null" json:"-"`
	BodyID       uint `gorm:"not null" json:"-"`
	CityID       uint `gorm:"not null" json:"-"`
	ModelID      uint `gorm:"not null" json:"-"`
	// Relationships with details
	User       *User       `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Category   *Category   `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Brand      *Brand      `gorm:"foreignKey:BrandID" json:"brand,omitempty"`
	Model      *Model      `gorm:"foreignKey:ModelID" json:"model,omitempty"`
	Generation *Generation `gorm:"foreignKey:GenerationID" json:"generation,omitempty"`
	City       *City       `gorm:"foreignKey:CityID" json:"city,omitempty"`
	Color      *Color      `gorm:"foreignKey:ColorID" json:"color,omitempty"`
	Body       *Body       `gorm:"foreignKey:BodyID" json:"body,omitempty"`
	// In place params
	Price            string `json:"price,omitempty"`
	EngineVolume     string `json:"engine_volume,omitempty"`
	Mileage          string `json:"mileage,omitempty"`
	CustomsClearance bool   `json:"customs_clearance,omitempty"`
	Description      string `json:"description,omitempty"`
	SteeringWheel    string `json:"steering_wheel,omitempty"` // Left/Right
	WheelDrive       string `json:"wheel_drive,omitempty"`    // FWD/RWD/AWD
	// Avatar url
	AvatarSource string `json:"avatar_source,omitempty"`
	// Time params
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
}
