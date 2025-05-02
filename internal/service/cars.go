package service

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/bwjson/kolesa_api/internal/adapter/http/handler/dto"
	"github.com/bwjson/kolesa_api/internal/models"
	"github.com/bwjson/kolesa_api/internal/repository"
	"github.com/bwjson/kolesa_api/pkg/s3"
	"strings"
)

type CarsService struct {
	repo        repository.Cars
	userRepo    repository.Users
	detailsRepo repository.Details
	s3          *s3.S3Client
}

func NewCarsService(repo repository.Cars, userRepo repository.Users, detailsRepo repository.Details, s3 *s3.S3Client) *CarsService {
	return &CarsService{repo: repo, userRepo: userRepo, s3: s3}
}

func (s *CarsService) Create(ctx context.Context, dto dto.CreateCarDTO) (int, error) {
	user, err := s.userRepo.GetByPhoneNumber(ctx, dto.CurrentUserPhoneNumber)
	if err != nil {
		return 0, err
	}

	car := models.Car{
		UserID:           uint(user.Id),
		CategoryID:       dto.CategoryID,
		BrandID:          dto.BrandID,
		ColorID:          dto.ColorID,
		GenerationID:     dto.GenerationID,
		BodyID:           dto.BodyID,
		CityID:           dto.CityID,
		ModelID:          dto.ModelID,
		Price:            dto.Price,
		EngineVolume:     dto.EngineVolume,
		Mileage:          dto.Mileage,
		CustomsClearance: dto.CustomsClearance,
		Description:      dto.Description,
		SteeringWheel:    dto.SteeringWheel,
		WheelDrive:       dto.WheelDrive,
	}

	var uploadedUrls []string

	for i, base64Image := range dto.Images {
		dataIdx := strings.Index(base64Image, "base64, ")
		if dataIdx == -1 {
			return 0, errors.New("Invalid base64 format")
		}

		fileBytes, err := base64.StdEncoding.DecodeString(base64Image[dataIdx+7:])
		if err != nil {
			return 0, errors.New("Cannot decode base64 image")
		}

		// logic of unique identifier
		fileName := fmt.Sprintf("user-%d-test-%d", user.Id, i)

		_, err = s.s3.UploadFile(fileName, fileBytes)
		if err != nil {
			return 0, err
		}

		uploadedUrls = append(uploadedUrls, fileName)
	}

	// Logic of adding photos and car_photos

	return s.repo.Create(ctx, car)
}

func (s *CarsService) GetAllExtended(ctx context.Context, limit, offset int) ([]models.Car, int, error) {
	return s.repo.GetAllCarsExtended(ctx, limit, offset)
}

func (s *CarsService) GetAll(ctx context.Context, filters map[string]interface{}) ([]models.Car, int64, error) {
	credentials, err := s.s3.GetS3Credentials()
	if err != nil {
		return nil, 0, err
	}

	return s.repo.GetAllCars(ctx, filters, credentials.AuthToken)
}

func (s *CarsService) GetById(ctx context.Context, id int) (models.Car, error) {
	return s.repo.GetCarById(ctx, id)
}

func (s *CarsService) UpdateById(ctx context.Context, id int, car models.Car) error {
	return s.repo.UpdateById(ctx, id, car)
}

func (s *CarsService) DeleteById(ctx context.Context, id int) error {
	return s.repo.DeleteById(ctx, id)
}

func (s *CarsService) SearchCars(ctx context.Context, query string, limit, offset int) ([]models.Car, int64, error) {
	credentials, err := s.s3.GetS3Credentials()
	if err != nil {
		return nil, 0, err
	}

	return s.repo.SearchCars(ctx, query, credentials.AuthToken, limit, offset)
}
