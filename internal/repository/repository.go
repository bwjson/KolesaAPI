package repository

import (
	"context"
	"github.com/bwjson/kolesa_api/internal/models"
	"gorm.io/gorm"
)

type Cars interface {
	Create(ctx context.Context, good models.Car) (int, error)
	GetAllCars(ctx context.Context, filters map[string]interface{}, authToken string) ([]models.Car, int64, error) // second param is total_count
	GetAllCarsExtended(ctx context.Context, limit, offset int) ([]models.Car, int, error)
	GetCarById(ctx context.Context, id int) (models.Car, error)
	UpdateById(ctx context.Context, id int, car models.Car) error
	DeleteById(ctx context.Context, id int) error
	SearchCars(ctx context.Context, query, authToken string, limit, offset int) ([]models.Car, int64, error)
}

type Details interface {
	GetAllCities(ctx context.Context) ([]models.City, error)
	GetAllBrands(ctx context.Context) ([]models.Brand, error)
	GetAllModels(ctx context.Context, brandSource string) ([]models.Model, error)
	GetAllGenerations(ctx context.Context, modelSource string) ([]models.Generation, error)
	GetAllCategories(ctx context.Context) ([]models.Category, error)
	GetAllBodies(ctx context.Context) ([]models.Body, error)
	GetAllColors(ctx context.Context) ([]models.Color, error)
	GetSourceById(ctx context.Context, carId int) (string, error)
	AddSourceUrl(ctx context.Context, photo models.CarPhoto) error
}

type Users interface {
	Create(ctx context.Context, user models.User) (int, error)
	GetAll(ctx context.Context) ([]models.User, error)
	GetByID(ctx context.Context, id int) (models.User, error)
	GetByPhoneNumber(ctx context.Context, phoneNumber string) (*models.User, error)
	Update(ctx context.Context, id int, user models.User) error
	Delete(ctx context.Context, id int) error
}

type Repos struct {
	Cars    Cars
	Details Details
	Users   Users
}

func NewRepos(db *gorm.DB) *Repos {
	return &Repos{
		Cars:    NewCarsRepo(db),
		Details: NewDetailsRepo(db),
		Users:   NewUsersRepo(db),
	}
}
