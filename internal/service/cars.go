package service

import (
	"context"
	"github.com/bwjson/kolesa_api/internal/dto"
	"github.com/bwjson/kolesa_api/internal/repository"
	"github.com/bwjson/kolesa_api/pkg"
)

type CarsService struct {
	repo repository.Cars
	s3   *pkg.S3Client
}

func NewCarsService(repo repository.Cars, s3 *pkg.S3Client) *CarsService {
	return &CarsService{repo: repo, s3: s3}
}

func (s *CarsService) Create(ctx context.Context, good dto.Car) (int, error) {
	return s.repo.Create(ctx, good)
}

func (s *CarsService) GetAllExtended(ctx context.Context, limit, offset int) ([]dto.Car, int, error) {
	return s.repo.GetAllCarsExtended(ctx, limit, offset)
}

func (s *CarsService) GetAll(ctx context.Context, filters map[string]interface{}) ([]dto.Car, int64, error) {
	credentials, err := s.s3.GetS3Credentials()
	if err != nil {
		return nil, 0, err
	}

	return s.repo.GetAllCars(ctx, filters, credentials.AuthToken)
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

func (s *CarsService) SearchCars(ctx context.Context, query string, limit, offset int) ([]dto.Car, int64, error) {
	credentials, err := s.s3.GetS3Credentials()
	if err != nil {
		return nil, 0, err
	}

	return s.repo.SearchCars(ctx, query, credentials.AuthToken, limit, offset)
}
