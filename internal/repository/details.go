package repository

import (
	"context"
	"errors"
	"github.com/bwjson/kolesa_api/internal/dto"
	"gorm.io/gorm"
)

type DetailsRepo struct {
	db *gorm.DB
}

func NewDetailsRepo(db *gorm.DB) *DetailsRepo {
	return &DetailsRepo{db: db}
}

func (r *DetailsRepo) GetAllCities(ctx context.Context) ([]dto.City, error) {
	var cities []dto.City

	res := r.db.WithContext(ctx).Find(&cities)

	if res.Error != nil {
		return nil, errors.New("No cities found")
	}

	return cities, nil
}

func (r *DetailsRepo) GetAllBrands(ctx context.Context) ([]dto.Brand, error) {
	var brands []dto.Brand

	res := r.db.WithContext(ctx).Find(&brands)

	if res.Error != nil {
		return nil, errors.New("No brands found")
	}

	return brands, nil
}

func (r *DetailsRepo) GetAllModels(ctx context.Context, brandSource string) ([]dto.Model, error) {
	var models []dto.Model

	query := r.db.WithContext(ctx).Model(&models)

	if brandSource != "" {
		query = query.Where("brand_source = ?", brandSource)
	}

	res := query.Find(&models)

	if res.Error != nil {
		return nil, errors.New("No models found")
	}

	return models, nil
}

func (r *DetailsRepo) GetAllGenerations(ctx context.Context, modelSource string) ([]dto.Generation, error) {
	var generations []dto.Generation

	query := r.db.WithContext(ctx).Model(&generations)

	if modelSource != "" {
		query = query.Where("model_source = ?", modelSource)
	}

	res := query.Find(&generations)

	if res.Error != nil {
		return nil, errors.New("No generations found")
	}

	return generations, nil
}

func (r *DetailsRepo) GetAllCategories(ctx context.Context) ([]dto.Category, error) {
	var categories []dto.Category

	res := r.db.WithContext(ctx).Find(&categories)

	if res.Error != nil {
		return nil, errors.New("No categories found")
	}

	return categories, nil
}

func (r *DetailsRepo) GetAllBodies(ctx context.Context) ([]dto.Body, error) {
	var bodies []dto.Body

	res := r.db.WithContext(ctx).Find(&bodies)

	if res.Error != nil {
		return nil, errors.New("No bodies found")
	}

	return bodies, nil
}

func (r *DetailsRepo) GetAllColors(ctx context.Context) ([]dto.Color, error) {
	var colors []dto.Color

	res := r.db.WithContext(ctx).Find(&colors)

	if res.Error != nil {
		return nil, errors.New("No colors found")
	}

	return colors, nil
}

func (r *DetailsRepo) GetSourceById(ctx context.Context, carId int) (string, error) {
	var carPhoto dto.CarPhoto

	res := r.db.WithContext(ctx).Find(&carPhoto, "car_id = ?", carId)

	if res.Error != nil {
		return "", errors.New("No sources found")
	}

	return carPhoto.PhotoUrl, nil
}

func (r *DetailsRepo) AddSourceUrl(ctx context.Context, photo dto.CarPhoto) error {
	res := r.db.WithContext(ctx).Create(&photo)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *DetailsRepo) GetModelsByBrand(ctx context.Context, brand string) ([]dto.Model, error) {
	var models []dto.Model

	res := r.db.WithContext(ctx).Find(&models, "brand = ?", brand)

	if res.Error != nil {
		return nil, errors.New("No sources found")
	}

	return models, nil
}
