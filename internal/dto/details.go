package dto

type Category struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"unique;not null" json:"name"`
}

type Brand struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"unique;not null" json:"name"`
}

type Color struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"unique;not null" json:"name"`
}

type Model struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	BrandID uint   `gorm:"not null; constraint:OnDelete:CASCADE;" json:"brand_id"`
	Name    string `gorm:"unique;not null" json:"name"`
}

type Generation struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	ModelID uint   `gorm:"not null; constraint:OnDelete:CASCADE;" json:"model_id"`
	Name    string `gorm:"not null" json:"name"`
}

type Body struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"unique;not null" json:"name"`
}

type City struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"unique;not null" json:"name"`
}

type CarPhoto struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	CarID    uint   `gorm:"not null; index; constraint:OnDelete:CASCADE;" json:"car_id"`
	PhotoUrl string `gorm:"not null" json:"photo_url"`
}
