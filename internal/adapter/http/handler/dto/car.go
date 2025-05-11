package dto

import (
	"github.com/bwjson/kolesa_api/internal/models"
	"time"
)

type CreateCarDTO struct {
	CategorySource   string   `json:"category_source"`
	BrandSource      string   `json:"brand_source"`
	ColorSource      string   `json:"color_source"`
	GenerationSource string   `json:"generation_source"`
	BodySource       string   `json:"body_source"`
	CitySource       string   `json:"city_source"`
	ModelSource      string   `json:"model_source"`
	PhoneNumber      string   `json:"phone"`
	Price            string   `json:"price,omitempty"`
	EngineVolume     string   `json:"engine_volume,omitempty"`
	Mileage          string   `json:"mileage,omitempty"`
	CustomsClearance bool     `json:"customs_clearance,omitempty"`
	Description      string   `json:"description,omitempty"`
	SteeringWheel    string   `json:"steering_wheel,omitempty"`
	WheelDrive       string   `json:"wheel_drive,omitempty"`
	Year             string   `json:"year,omitempty"`
	Images           []string `json:"images,omitempty"`
}

type CarResponseDTO struct {
	// Relationships IDs
	ID uint `gorm:"primaryKey" json:"id"`
	// Relationships with details
	User       *SafeUserDTO       `json:"user,omitempty"`
	Category   *models.Category   `json:"category,omitempty"`
	Brand      *models.Brand      `json:"brand,omitempty"`
	Model      *models.Model      `json:"model,omitempty"`
	Generation *models.Generation `json:"generation,omitempty"`
	City       *models.City       `json:"city,omitempty"`
	Color      *models.Color      `json:"color,omitempty"`
	Body       *models.Body       `json:"body,omitempty"`
	// In place params
	Price            string `json:"price,omitempty"`
	EngineVolume     string `json:"engine_volume,omitempty"`
	Mileage          string `json:"mileage,omitempty"`
	CustomsClearance bool   `json:"customs_clearance,omitempty"`
	Description      string `json:"description,omitempty"`
	SteeringWheel    string `json:"steering_wheel,omitempty"` // Left/Right
	WheelDrive       string `json:"wheel_drive,omitempty"`    // FWD/RWD/AWD
	Year             string `json:"year,omitempty"`
	// Avatar url
	AvatarSource string `json:"avatar_source,omitempty"`
	// Time params
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
}
