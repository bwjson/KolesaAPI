package repository

import (
	"context"
	"errors"
	"github.com/bwjson/kolesa_api/internal/models"
	"gorm.io/gorm"
	"log"
)

type DetailsRepo struct {
	db *gorm.DB
}

func NewDetailsRepo(db *gorm.DB) *DetailsRepo {
	return &DetailsRepo{db: db}
}

func (r *DetailsRepo) GetAllCities(ctx context.Context) ([]models.City, error) {
	var cities []models.City

	res := r.db.WithContext(ctx).Find(&cities)

	if res.Error != nil {
		return nil, errors.New("No cities found")
	}

	return cities, nil
}

func (r *DetailsRepo) GetAllBrands(ctx context.Context) ([]models.Brand, error) {
	var brands []models.Brand

	res := r.db.WithContext(ctx).Find(&brands)

	if res.Error != nil {
		return nil, errors.New("No brands found")
	}

	return brands, nil
}

func (r *DetailsRepo) GetAllModels(ctx context.Context, brandSource string) ([]models.Model, error) {
	var models []models.Model

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

func (r *DetailsRepo) GetAllGenerations(ctx context.Context, modelSource string) ([]models.Generation, error) {
	var generations []models.Generation

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

func (r *DetailsRepo) GetAllCategories(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category

	res := r.db.WithContext(ctx).Find(&categories)

	if res.Error != nil {
		return nil, errors.New("No categories found")
	}

	return categories, nil
}

func (r *DetailsRepo) GetAllBodies(ctx context.Context) ([]models.Body, error) {
	var bodies []models.Body

	res := r.db.WithContext(ctx).Find(&bodies)

	if res.Error != nil {
		return nil, errors.New("No bodies found")
	}

	return bodies, nil
}

func (r *DetailsRepo) GetAllColors(ctx context.Context) ([]models.Color, error) {
	var colors []models.Color

	res := r.db.WithContext(ctx).Find(&colors)

	if res.Error != nil {
		return nil, errors.New("No colors found")
	}

	return colors, nil
}

func (r *DetailsRepo) GetSourceById(ctx context.Context, carId int) (string, error) {
	var carPhoto models.CarPhoto

	res := r.db.WithContext(ctx).Find(&carPhoto, "car_id = ?", carId)

	if res.Error != nil {
		return "", errors.New("No sources found")
	}

	return carPhoto.PhotoUrl, nil
}

func (r *DetailsRepo) AddSourceUrl(ctx context.Context, photo models.CarPhoto) error {
	res := r.db.WithContext(ctx).Create(&photo)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *DetailsRepo) GetModelsByBrand(ctx context.Context, brand string) ([]models.Model, error) {
	var models []models.Model

	res := r.db.WithContext(ctx).Find(&models, "brand = ?", brand)

	if res.Error != nil {
		return nil, errors.New("No sources found")
	}

	return models, nil
}

func (r *DetailsRepo) GetCategoryBySource(ctx context.Context, source string) (models.Category, error) {
	log.Println("here")
	var category models.Category

	res := r.db.WithContext(ctx).First(&category, "source = ?", source)

	if res.Error != nil {
		log.Println(res.Error)
		return models.Category{}, errors.New("No categories found")
	}

	return category, nil
}

func (r *DetailsRepo) GetBrandBySource(ctx context.Context, source string) (models.Brand, error) {
	var brand models.Brand
	res := r.db.WithContext(ctx).First(&brand, "source = ?", source)
	if res.Error != nil {
		return models.Brand{}, errors.New("no brand found")
	}
	return brand, nil
}

func (r *DetailsRepo) GetModelBySource(ctx context.Context, source string) (models.Model, error) {
	var model models.Model
	res := r.db.WithContext(ctx).First(&model, "source = ?", source)
	if res.Error != nil {
		return models.Model{}, errors.New("no model found")
	}
	return model, nil
}

func (r *DetailsRepo) GetColorBySource(ctx context.Context, source string) (models.Color, error) {
	var color models.Color
	res := r.db.WithContext(ctx).First(&color, "source = ?", source)
	if res.Error != nil {
		return models.Color{}, errors.New("no color found")
	}
	return color, nil
}

func (r *DetailsRepo) GetBodyBySource(ctx context.Context, source string) (models.Body, error) {
	var body models.Body
	res := r.db.WithContext(ctx).First(&body, "source = ?", source)
	if res.Error != nil {
		return models.Body{}, errors.New("no body found")
	}
	return body, nil
}

func (r *DetailsRepo) GetGenerationBySource(ctx context.Context, source string) (models.Generation, error) {
	var generation models.Generation
	res := r.db.WithContext(ctx).First(&generation, "source = ?", source)
	if res.Error != nil {
		return models.Generation{}, errors.New("no generation found")
	}
	return generation, nil
}

func (r *DetailsRepo) GetCityBySource(ctx context.Context, source string) (models.City, error) {
	var city models.City
	res := r.db.WithContext(ctx).First(&city, "source = ?", source)
	if res.Error != nil {
		return models.City{}, errors.New("no city found")
	}
	return city, nil
}
