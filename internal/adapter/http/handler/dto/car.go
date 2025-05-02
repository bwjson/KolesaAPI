package dto

type CreateCarDTO struct {
	CategoryID             uint     `json:"category_id"`
	BrandID                uint     `json:"brand_id"`
	ColorID                uint     `json:"color_id"`
	GenerationID           uint     `json:"generation_id"`
	BodyID                 uint     `json:"body_id"`
	CityID                 uint     `json:"city_id"`
	ModelID                uint     `json:"model_id"`
	CurrentUserPhoneNumber string   `json:"current_user_phone_number"`
	Price                  string   `json:"price,omitempty"`
	EngineVolume           string   `json:"engine_volume,omitempty"`
	Mileage                string   `json:"mileage,omitempty"`
	CustomsClearance       bool     `json:"customs_clearance,omitempty"`
	Description            string   `json:"description,omitempty"`
	SteeringWheel          string   `json:"steering_wheel,omitempty"`
	WheelDrive             string   `json:"wheel_drive,omitempty"`
	Images                 []string `json:"images,omitempty"`
}
