package dto

type CreateCarDTO struct {
	CategorySource   string   `json:"category_source"`
	BrandSource      string   `json:"brand_source"`
	ColorSource      string   `json:"color_source"`
	GenerationSource string   `json:"generation_source"`
	BodySource       string   `json:"body_source"`
	CitySource       string   `json:"city_source"`
	ModelSource      string   `json:"model_source"`
	PhoneNumber      string   `json:"phone_number"`
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
