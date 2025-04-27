package service

import (
	"context"
	"github.com/bwjson/kolesa_api/internal/dto"
	"github.com/bwjson/kolesa_api/internal/repository"
	"github.com/bwjson/kolesa_api/pkg"
)

type Cars interface {
	Create(ctx context.Context, good dto.Car) (int, error)
	GetAllExtended(ctx context.Context, limit, offset int) ([]dto.Car, int, error)
	GetById(ctx context.Context, id int) (dto.Car, error)
	GetAll(ctx context.Context, filters map[string]interface{}) ([]dto.Car, int64, error) // second param is total_count
	UpdateById(ctx context.Context, id int, good dto.Car) error
	DeleteById(ctx context.Context, id int) error
	SearchCars(ctx context.Context, query string, limit, offset int) ([]dto.Car, int64, error)
}

type Services struct {
	Cars Cars
	S3   *pkg.S3Client
}

func NewServices(repos *repository.Repos, s3 *pkg.S3Client) *Services {
	return &Services{
		Cars: NewCarsService(repos.Cars, s3),
	}
}
