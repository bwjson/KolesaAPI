package service

import (
	"context"
	"github.com/bwjson/kolesa_api/internal/adapter/http/handler/dto"
	"github.com/bwjson/kolesa_api/internal/models"
	"github.com/bwjson/kolesa_api/internal/repository"
	"github.com/bwjson/kolesa_api/pkg/s3"
)

type Cars interface {
	Create(ctx context.Context, carDTO dto.CreateCarDTO) (uint, error)
	GetAllExtended(ctx context.Context, limit, offset int) ([]models.Car, int, error)
	GetById(ctx context.Context, id int) (models.Car, error)
	GetAll(ctx context.Context, filters map[string]interface{}) ([]models.Car, int64, error) // second param is total_count
	UpdateById(ctx context.Context, id int, good models.Car) error
	DeleteById(ctx context.Context, id int) error
	SearchCars(ctx context.Context, query string, limit, offset int) ([]models.Car, int64, error)
}

// divide to different parts e.g. cities, brands in future for convience
type Details interface {
	GetAllBrands(ctx context.Context) ([]models.Brand, error)
	GetAllModels(ctx context.Context, brandSource string) ([]models.Model, error)
	GetAllGenerations(ctx context.Context, modelSource string) ([]models.Generation, error)

	GetAllCities(ctx context.Context) ([]models.City, error)
	GetAllCategories(ctx context.Context) ([]models.Category, error)
	GetAllBodies(ctx context.Context) ([]models.Body, error)
	GetAllColors(ctx context.Context) ([]models.Color, error)
	GetSourceById(ctx context.Context, carId int) (string, error)

	GetCategoryBySource(ctx context.Context, source string) (models.Category, error)
	GetBrandBySource(ctx context.Context, source string) (models.Brand, error)
	GetModelBySource(ctx context.Context, source string) (models.Model, error)
	GetColorBySource(ctx context.Context, source string) (models.Color, error)
	GetBodyBySource(ctx context.Context, source string) (models.Body, error)
	GetGenerationBySource(ctx context.Context, source string) (models.Generation, error)
	GetCityBySource(ctx context.Context, source string) (models.City, error)

	AddSourceUrl(ctx context.Context, photo models.CarPhoto) error
}

type Users interface {
	GetAll(ctx context.Context) ([]models.User, error)
	GetByID(ctx context.Context, id int) (models.User, error)
	GetByPhoneNumber(ctx context.Context, phoneNumber string) (*models.User, error)
	Create(ctx context.Context, user models.User) (int, error)
	Update(ctx context.Context, id int, user models.User) error
	Delete(ctx context.Context, id int) error
}

type Services struct {
	Cars    Cars
	Details Details
	Users   Users
	S3      *s3.S3Client
}

func NewServices(repos *repository.Repos, s3 *s3.S3Client) *Services {
	return &Services{
		Cars:    NewCarsService(repos.Cars, repos.Users, repos.Details, s3),
		Details: NewDetailsService(repos.Details),
		Users:   NewUsersService(repos.Users),
	}
}
