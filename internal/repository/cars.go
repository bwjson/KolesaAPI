package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/bwjson/kolesa_api/internal/models"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
	"log"
)

type CarsRepo struct {
	db *gorm.DB
}

func NewCarsRepo(db *gorm.DB) *CarsRepo {
	return &CarsRepo{db: db}
}

func (r *CarsRepo) SearchCars(ctx context.Context, query, authToken string, limit, offset int) ([]models.Car, int64, error) {
	var cars []models.Car
	var totalCount int64

	sql_query := `
		to_tsvector('simple', coalesce(cars.description, '') || ' ' || 
			coalesce(brands.name, '') || ' ' ||
			coalesce(models.name, '') || ' ' ||
			coalesce(categories.name, '') || ' ' ||
			coalesce(generations.name, '') || ' ' ||
			coalesce(cities.name, '') || ' ' ||
			coalesce(colors.name, '') || ' ' ||
			coalesce(bodies.name, '')
		) @@ plainto_tsquery('simple', ?)
	`

	base_query := r.db.Model(&models.Car{}).
		Joins("LEFT JOIN brands ON brands.id = cars.brand_id").
		Joins("LEFT JOIN models ON models.id = cars.model_id").
		Joins("LEFT JOIN categories ON categories.id = cars.category_id").
		Joins("LEFT JOIN generations ON generations.id = cars.generation_id").
		Joins("LEFT JOIN cities ON cities.id = cars.city_id").
		Joins("LEFT JOIN colors ON colors.id = cars.color_id").
		Joins("LEFT JOIN bodies ON bodies.id = cars.body_id").
		Where(sql_query, query).
		Preload("Brand").
		Preload("Model").
		Preload("Category").
		Preload("Generation").
		Preload("City").
		Preload("Color").
		Preload("Body")

	err := base_query.Count(&totalCount).Error
	if err != nil {
		return nil, 0, errors.New("Could not calculate the totalCount param")
	}

	// Offset means pages not elements
	offset *= limit

	log.Println(offset)

	err = base_query.Limit(limit).Offset(offset).Find(&cars).Error
	if err != nil {
		return nil, 0, errors.New("Could not get the data")
	}

	// Authorization token
	for i := range cars {
		cars[i].AvatarSource += "?Authorization=" + authToken
	}

	return cars, totalCount, nil
}

func (r *CarsRepo) Create(ctx context.Context, good models.Car) (int, error) {

	return 0, nil
}

func (r *CarsRepo) GetAllCarsExtended(ctx context.Context, limit, offset int) ([]models.Car, int, error) {
	var cars []models.Car

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

func (r *CarsRepo) GetAllCars(ctx context.Context, filters map[string]interface{}, authToken string) ([]models.Car, int64, error) {
	var cars []models.Car
	var totalCount int64
	var limit, offset int

	baseQuery := r.db.WithContext(ctx).Model(&models.Car{})

	// Filters
	if v, ok := filters["categorySource"].(string); ok && v != "" {
		baseQuery.
			Joins("JOIN categories on categories.id = cars.category_id").
			Where("LOWER(categories.source) = LOWER(?)", v)
	}

	if v, ok := filters["brandSource"].(string); ok && v != "" {
		baseQuery.
			Joins("JOIN brands on brands.id = cars.brand_id").
			Where("LOWER(brands.source) = LOWER (?)", v)
	}

	if v, ok := filters["modelSource"].(string); ok && v != "" {
		baseQuery = baseQuery.Joins("JOIN models on models.id = cars.model_id").
			Where("LOWER(models.source) = LOWER(?)", v)
	}

	if v, ok := filters["generationSource"].(string); ok && v != "" {
		baseQuery = baseQuery.Joins("JOIN generations on generations.id = cars.generation_id").
			Where("LOWER(generations.source) = LOWER(?)", v)
	}

	if v, ok := filters["citySource"].(string); ok && v != "" {
		baseQuery = baseQuery.Joins("JOIN cities on cities.id = cars.city_id").
			Where("LOWER(cities.source) = LOWER(?)", v)
	}

	if v, ok := filters["colorSource"].(string); ok && v != "" {
		baseQuery = baseQuery.Joins("JOIN colors on colors.id = cars.color_id").
			Where("LOWER(colors.source) = LOWER(?)", v)
	}

	if v, ok := filters["bodySource"].(string); ok && v != "" {
		baseQuery = baseQuery.Joins("JOIN bodies on bodies.id = cars.body_id").
			Where("LOWER(bodies.source) = LOWER(?)", v)
	}

	if v, ok := filters["steeringWheel"].(string); ok && v != "" {
		baseQuery = baseQuery.Where("LOWER(cars.steering_wheel) = LOWER(?)", v)
	}

	if v, ok := filters["wheelDrive"].(string); ok && v != "" {
		baseQuery = baseQuery.Where("LOWER(cars.wheel_drive) = LOWER(?)", v)
	}

	// Between params
	if vStart, ok := filters["priceStart"].(int); ok && vStart > 0 {
		baseQuery = baseQuery.Where("CAST(cars.price AS INTEGER) >= ?", vStart)
	}
	if vEnd, ok := filters["priceEnd"].(int); ok && vEnd > 0 {
		baseQuery = baseQuery.Where("CAST(cars.price AS INTEGER) <= ?", vEnd)
	}

	if vStart, ok := filters["engineStart"].(float64); ok && vStart > 0 {
		baseQuery = baseQuery.Where("CAST(cars.engine_volume AS FLOAT) >= ?", vStart)
	}
	if vEnd, ok := filters["engineEnd"].(float64); ok && vEnd > 0 {
		baseQuery = baseQuery.Where("CAST(cars.engine_volume AS FLOAT) <= ?", vEnd)
	}

	if vStart, ok := filters["mileageStart"].(int); ok && vStart >= 0 {
		baseQuery = baseQuery.Where("CAST(cars.mileage AS INTEGER) >= ?", vStart)
	}
	if vEnd, ok := filters["mileageEnd"].(int); ok && vEnd > 0 {
		baseQuery = baseQuery.Where("CAST(cars.mileage AS INTEGER) <= ?", vEnd)
	}

	// Get total_count before pagination params and exclude offset
	err := baseQuery.Count(&totalCount).Error
	if err != nil {
		return nil, 0, errors.New("Failed to count cars")
	}

	// Pagination
	if v, ok := filters["limit"].(int); ok {
		limit = v
	}

	if v, ok := filters["offset"].(int); ok {
		offset = v
	}

	// From client we are getting number of pages
	offset *= limit

	err = baseQuery.
		Select("cars.id", "cars.price", "cars.category_id", "cars.brand_id", "cars.model_id", "cars.avatar_source").
		Preload("Category").
		Preload("Brand").
		Preload("Model").
		Limit(limit).Offset(offset).Find(&cars).Error
	if err != nil {
		return nil, 0, errors.New("No cars found")
	}

	// Authorization token
	for i := range cars {
		cars[i].AvatarSource += "?Authorization=" + authToken
	}

	return cars, totalCount, nil
}

func (r *CarsRepo) GetCarById(ctx context.Context, id int) (models.Car, error) {
	var car models.Car

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
			return models.Car{}, fmt.Errorf("Car with ID = %d not found", id)
		}
		return models.Car{}, fmt.Errorf("Database error: %w", result.Error)
	}

	return car, nil
}

func (r *CarsRepo) UpdateById(ctx context.Context, id int, good models.Car) error {
	return nil
}

func (r *CarsRepo) DeleteById(ctx context.Context, id int) error {
	return nil
}
