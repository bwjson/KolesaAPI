package service

import (
	"context"
	"github.com/bwjson/kolesa_api/internal/dto"
	"github.com/bwjson/kolesa_api/internal/repository"
)

type Cars interface {
	Create(ctx context.Context, good dto.Car) (int, error)
	GetAllExtended(ctx context.Context, limit, offset int) ([]dto.Car, int, error)
	GetById(ctx context.Context, id int) (dto.Car, error)
	GetAll(ctx context.Context, limit, offset int) ([]dto.Car, int, error) // second param is total_count
	UpdateById(ctx context.Context, id int, good dto.Car) error
	DeleteById(ctx context.Context, id int) error
}

type Services struct {
	Cars Cars
}

func NewServices(repos *repository.Repos) *Services {
	return &Services{Cars: NewCarsService(repos.Cars)}
}
