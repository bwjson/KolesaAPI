package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/bwjson/kolesa_api/internal/dto"
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

func (r *CarsRepo) SearchCars(ctx context.Context, query string) ([]dto.Car, error) {
	var cars []dto.Car

	searchVector := `
		to_tsvector('simple',
			coalesce(price, '') || ' ' ||
			coalesce(engine_volume, '') || ' ' ||
			coalesce(mileage, '') || ' ' ||
			coalesce(description, '') || ' ' ||
			coalesce(steering_wheel, '') || ' ' ||
			coalesce(wheel_drive, '')
		) @@ plainto_tsquery('simple', ?)
	`

	log.Println(searchVector)

	res := r.db.
		WithContext(ctx).
		Model(dto.Car{}).
		Find(&cars)

	log.Println(query)

	if res.Error != nil {
		return nil, res.Error
	}
	return cars, nil
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

func (r *CarsRepo) GetAllCars(ctx context.Context, filters map[string]interface{}, authToken string) ([]dto.Car, int64, error) {
	var cars []dto.Car
	var totalCount int64
	var limit, offset int

	baseQuery := r.db.WithContext(ctx).Model(&dto.Car{})

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

	log.Println(totalCount, offset)

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
