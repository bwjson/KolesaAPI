package repository

import (
	"context"
	"github.com/bwjson/kolesa_api/internal/dto"
	"gorm.io/gorm"
)

type Cars interface {
	Create(ctx context.Context, good dto.Car) (int, error)
	GetAllCars(ctx context.Context, limit, offset int) ([]dto.Car, int, error)         // second param is total_count
	GetAllCarsExtended(ctx context.Context, limit, offset int) ([]dto.Car, int, error) // second param is total_count
	GetCarById(ctx context.Context, id int) (dto.Car, error)
	UpdateById(ctx context.Context, id int, car dto.Car) error
	DeleteById(ctx context.Context, id int) error
}

type Details interface {
	GetAllCities(ctx context.Context) ([]dto.City, error)
	GetAllBrands(ctx context.Context) ([]dto.Brand, error)
	GetAllModels(ctx context.Context) ([]dto.Model, error)
	GetAllCategories(ctx context.Context) ([]dto.Category, error)
	GetAllBodies(ctx context.Context) ([]dto.Body, error)
	GetAllGenerations(ctx context.Context) ([]dto.Generation, error)
	GetAllColors(ctx context.Context) ([]dto.Color, error)
	GetSourceById(ctx context.Context, carId int) (string, error)
}

type Repos struct {
	Cars    Cars
	Details Details
}

func NewRepos(db *gorm.DB) *Repos {
	return &Repos{
		Cars:    NewCarsRepo(db),
		Details: NewDetailsRepo(db),
	}
}
