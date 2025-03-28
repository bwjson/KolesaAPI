package service

import (
	"context"
	"github.com/bwjson/kolesa_api/internal/dto"
	"github.com/bwjson/kolesa_api/internal/repository"
)

type CarsService struct {
	repo repository.Cars
}

func NewCarsService(repo repository.Cars) *CarsService {
	return &CarsService{repo: repo}
}

func (s *CarsService) Create(ctx context.Context, good dto.Car) (int, error) {
	return s.repo.Create(ctx, good)
}

func (s *CarsService) GetAllExtended(ctx context.Context, limit, offset int) ([]dto.Car, int, error) {
	return s.repo.GetAllCarsExtended(ctx, limit, offset)
}

func (s *CarsService) GetAll(ctx context.Context, limit, offset int) ([]dto.Car, int, error) {
	return s.repo.GetAllCars(ctx, limit, offset)
}

func (s *CarsService) GetById(ctx context.Context, id int) (dto.Car, error) {
	return s.repo.GetCarById(ctx, id)
}

func (s *CarsService) UpdateById(ctx context.Context, id int, car dto.Car) error {
	return s.repo.UpdateById(ctx, id, car)
}

func (s *CarsService) DeleteById(ctx context.Context, id int) error {
	return s.repo.DeleteById(ctx, id)
}
