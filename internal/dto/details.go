package dto

// Filters embedded

type Category struct {
	ID     uint   `gorm:"primaryKey" json:"-"`
	Name   string `gorm:"unique;not null" json:"name"`
	Source string `json:"source"`
}

type Brand struct {
	ID     uint   `gorm:"primaryKey" json:"-"`
	Name   string `gorm:"unique;not null" json:"name"`
	Source string `json:"source"` // add later gorm tags
}

type Model struct {
	ID          uint   `gorm:"primaryKey" json:"-"`
	BrandID     uint   `gorm:"not null; constraint:OnDelete:CASCADE;" json:"-"`
	BrandSource string `json:"brand_source"`
	Name        string `gorm:"unique;not null" json:"name"`
	Source      string `json:"source"` // add later gorm tags
}

type Generation struct {
	ID          uint   `gorm:"primaryKey" json:"-"`
	ModelID     uint   `gorm:"not null; constraint:OnDelete:CASCADE;" json:"-"`
	ModelSource string `json:"model_source"`
	Name        string `gorm:"not null" json:"name"`
	Source      string `json:"source"`
}

type Color struct {
	ID     uint   `gorm:"primaryKey" json:"-"`
	Name   string `gorm:"unique;not null" json:"name"`
	Source string `json:"source"`
}

type Body struct {
	ID     uint   `gorm:"primaryKey" json:"-"`
	Name   string `gorm:"unique;not null" json:"name"`
	Source string `json:"source"`
}

type City struct {
	ID     uint   `gorm:"primaryKey" json:"-"`
	Name   string `gorm:"unique;not null" json:"name"`
	Source string `json:"source"`
}

type CarPhoto struct {
	ID       uint   `gorm:"primaryKey" json:"-"`
	CarID    uint   `gorm:"not null; index; constraint:OnDelete:CASCADE;" json:"-"`
	PhotoUrl string `gorm:"not null" json:"photo_url"`
}
