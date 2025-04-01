package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/bwjson/kolesa_api/internal/dto"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type CarsRepo struct {
	db *gorm.DB
}

func NewCarsRepo(db *gorm.DB) *CarsRepo {
	return &CarsRepo{db: db}
}

func (r *CarsRepo) Create(ctx context.Context, good dto.Car) (int, error) {
	//var id int
	//
	//err := r.db.QueryRow("INSERT INTO goods (name, description, photo_url, price) VALUES ($1, $2, $3, $4) RETURNING id",
	//	good.Name, good.Description, good.PhotoUrl, good.Price).Scan(&id)
	//
	//if err != nil {
	//	return 0, err
	//}
	//
	return 0, nil
}

func (r *CarsRepo) GetAllCarsExtended(ctx context.Context, limit, offset int) ([]dto.Car, int, error) {
	var cars []dto.Car

	res := r.db.WithContext(ctx).
		Preload("User").
		Preload("Category").
		Preload("Brand").
		Preload("Color").
		Preload("Generation").
		Preload("Body").
		Preload("City").
		Preload("Model").
		Limit(limit).
		Offset(offset).
		Find(&cars)

	if res.Error != nil {
		return nil, 0, errors.New("No cars found")
	}

	return cars, int(res.RowsAffected), nil
}

func (r *CarsRepo) GetAllCars(ctx context.Context, limit, offset int,
	brandSource, modelSource, generationSource, citySource string, authToken string) ([]dto.Car, int, error) {
	var cars []dto.Car

	query := r.db.WithContext(ctx).
		Select("cars.id", "cars.price",
			"cars.category_id", "cars.brand_id",
			"cars.model_id", "cars.avatar_source").
		Preload("Category").
		Preload("Brand").
		Preload("Model")

	if brandSource != "" {
		query.
			Joins("JOIN brands on brands.id = cars.brand_id").
			Where("LOWER(brands.source) = LOWER (?)", brandSource)
	}

	if modelSource != "" {
		query.
			Joins("JOIN models on models.id = cars.model_id").
			Where("LOWER(models.source) = LOWER(?)", modelSource)
	}

	if generationSource != "" {
		query.
			Joins("JOIN generations on generations.id = cars.generation_id").
			Where("LOWER(generations.source) = LOWER(?)", generationSource)
	}

	if citySource != "" {
		query.
			Joins("JOIN cities on cities.id = cars.city_id").
			Where("LOWER(cities.source) = LOWER(?)", citySource)
	}

	res := query.Limit(limit).Offset(offset).Find(&cars)

	// Authorization token
	for i := range cars {
		cars[i].AvatarSource += "?Authorization=" + authToken
	}

	if res.Error != nil {
		return nil, 0, errors.New("No cars found")
	}

	return cars, int(res.RowsAffected), nil
}

func (r *CarsRepo) GetCarById(ctx context.Context, id int) (dto.Car, error) {
	var car dto.Car

	result := r.db.WithContext(ctx).
		Preload("User").
		Preload("Category").
		Preload("Brand").
		Preload("Color").
		Preload("Generation").
		Preload("Body").
		Preload("City").First(&car, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return dto.Car{}, fmt.Errorf("Car with ID = %d not found", id)
		}
		return dto.Car{}, fmt.Errorf("Database error: %w", result.Error)
	}

	return car, nil
}

func (r *CarsRepo) UpdateById(ctx context.Context, id int, good dto.Car) error {
	return nil
}

func (r *CarsRepo) DeleteById(ctx context.Context, id int) error {
	return nil
}
