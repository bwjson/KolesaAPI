package repository

import (
	"context"
	"github.com/bwjson/kolesa_api/internal/dto"
	"gorm.io/gorm"
)

type Cars interface {
	Create(ctx context.Context, good dto.Car) (int, error)
	GetAll(ctx context.Context, limit, offset int) ([]dto.Car, int, error)         // second param is total_count
	GetAllExtended(ctx context.Context, limit, offset int) ([]dto.Car, int, error) // second param is total_count
	GetById(ctx context.Context, id int) (dto.Car, error)
	UpdateById(ctx context.Context, id int, car dto.Car) error
	DeleteById(ctx context.Context, id int) error
}

type Repos struct {
	Cars Cars
}

func NewRepos(db *gorm.DB) *Repos {
	return &Repos{
		Cars: NewCarsRepo(db)}
}
