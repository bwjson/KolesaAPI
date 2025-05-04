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
	"gorm.io/gorm"
	"log"
)

type CarsService struct {
	repo        repository.Cars
	userRepo    repository.Users
	detailsRepo repository.Details
	s3          *s3.S3Client
}

func NewCarsService(repo repository.Cars, userRepo repository.Users, detailsRepo repository.Details, s3 *s3.S3Client) *CarsService {
	return &CarsService{repo: repo, userRepo: userRepo, detailsRepo: detailsRepo, s3: s3}
}

func (s *CarsService) Create(ctx context.Context, dto dto.CreateCarDTO) (uint, error) {
	user, err := s.userRepo.GetByPhoneNumber(ctx, dto.PhoneNumber)
	if err != nil {
		return 0, err
	}

	category, err := s.detailsRepo.GetCategoryBySource(ctx, dto.CategorySource)
	if err != nil {
		return 0, err
	}

	brand, err := s.detailsRepo.GetBrandBySource(ctx, dto.BrandSource)
	if err != nil {
		return 0, err
	}

	model, err := s.detailsRepo.GetModelBySource(ctx, dto.ModelSource)
	if err != nil {
		return 0, err
	}

	generation, err := s.detailsRepo.GetGenerationBySource(ctx, dto.GenerationSource)
	if err != nil {
		return 0, err
	}

	body, err := s.detailsRepo.GetBodyBySource(ctx, dto.BodySource)
	if err != nil {
		return 0, err
	}

	color, err := s.detailsRepo.GetColorBySource(ctx, dto.ColorSource)
	if err != nil {
		return 0, err
	}

	city, err := s.detailsRepo.GetCityBySource(ctx, dto.CitySource)
	if err != nil {
		return 0, err
	}

	car := models.Car{
		UserID:           uint(user.Id),
		CategoryID:       category.ID,
		BrandID:          brand.ID,
		ModelID:          model.ID,
		GenerationID:     generation.ID,
		BodyID:           body.ID,
		ColorID:          color.ID,
		CityID:           city.ID,
		Price:            dto.Price,
		EngineVolume:     dto.EngineVolume,
		Mileage:          dto.Mileage,
		CustomsClearance: dto.CustomsClearance,
		Description:      dto.Description,
		SteeringWheel:    dto.SteeringWheel,
		WheelDrive:       dto.WheelDrive,
	}

	log.Println(car.UserID)

	// Transaction if other part doesnt work so we can rollback the changes
	carId, err := s.repo.Create(ctx, car)
	if err != nil {
		return 0, err
	}

	uniqueFilenamePart := fmt.Sprintf("%d-%d", carId, user.Id)

	var uploadedUrls []string

	for i, base64Image := range dto.Images {
		fileBytes, err := base64.StdEncoding.DecodeString(base64Image)
		if err != nil {
			return 0, errors.New("Cannot decode base64 image")
		}

		fileName := uniqueFilenamePart + fmt.Sprintf("-%d", i)

		_, err = s.s3.UploadFile(fileName, fileBytes)
		if err != nil {
			return 0, err
		}

		uploadedUrls = append(uploadedUrls, fileName)
	}

	// avatar change this with config variables
	err = s.repo.UpdateField(ctx, int(carId), "avatar_source",
		fmt.Sprintf("https://f006.backblazeb2.com/file/kolesa/main_photos/%v.jpg", uploadedUrls[0]))
	if err != nil {
		return 0, err
	}

	// adding to car_photos
	for _, photoUrl := range uploadedUrls {
		car_photo := models.CarPhoto{
			CarID:    carId,
			PhotoUrl: photoUrl,
		}
		err = s.detailsRepo.AddSourceUrl(ctx, car_photo)
	}

	return carId, nil
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
	car, err := s.repo.GetCarById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return car, ErrNotFound
		}
		return car, err
	}

	return car, nil
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
