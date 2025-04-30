package service

import (
	"context"
	"github.com/bwjson/kolesa_api/internal/adapter/http/handler/dto"
	"github.com/bwjson/kolesa_api/internal/models"
	"github.com/bwjson/kolesa_api/internal/repository"
	"github.com/bwjson/kolesa_api/pkg/s3"
)

type Cars interface {
	Create(ctx context.Context, carDTO dto.CreateCarDTO) (int, error)
	GetAllExtended(ctx context.Context, limit, offset int) ([]models.Car, int, error)
	GetById(ctx context.Context, id int) (models.Car, error)
	GetAll(ctx context.Context, filters map[string]interface{}) ([]models.Car, int64, error) // second param is total_count
	UpdateById(ctx context.Context, id int, good models.Car) error
	DeleteById(ctx context.Context, id int) error
	SearchCars(ctx context.Context, query string, limit, offset int) ([]models.Car, int64, error)
}

type Services struct {
	Cars Cars
	S3   *s3.S3Client
}

func NewServices(repos *repository.Repos, s3 *s3.S3Client) *Services {
	return &Services{
		Cars: NewCarsService(repos.Cars, repos.Users, s3),
	}
}
