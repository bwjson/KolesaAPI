package service

import (
	"context"
	"github.com/bwjson/kolesa_api/internal/models"
	"github.com/bwjson/kolesa_api/internal/repository"
)

type DetailsService struct {
	repo repository.Details
}

func NewDetailsService(repo repository.Details) *DetailsService {
	return &DetailsService{repo: repo}
}

func (s *DetailsService) GetAllCities(ctx context.Context) ([]models.City, error) {
	return s.repo.GetAllCities(ctx)
}

func (s *DetailsService) GetAllBrands(ctx context.Context) ([]models.Brand, error) {
	return s.repo.GetAllBrands(ctx)
}

func (s *DetailsService) GetAllModels(ctx context.Context, brandSource string) ([]models.Model, error) {
	return s.repo.GetAllModels(ctx, brandSource)
}

func (s *DetailsService) GetAllGenerations(ctx context.Context, modelSource string) ([]models.Generation, error) {
	return s.repo.GetAllGenerations(ctx, modelSource)
}

func (s *DetailsService) GetAllCategories(ctx context.Context) ([]models.Category, error) {
	return s.repo.GetAllCategories(ctx)
}

func (s *DetailsService) GetAllBodies(ctx context.Context) ([]models.Body, error) {
	return s.repo.GetAllBodies(ctx)
}

func (s *DetailsService) GetAllColors(ctx context.Context) ([]models.Color, error) {
	return s.repo.GetAllColors(ctx)
}

func (s *DetailsService) GetSourceById(ctx context.Context, carId int) (string, error) {
	return s.repo.GetSourceById(ctx, carId)
}

func (s *DetailsService) GetCategoryBySource(ctx context.Context, source string) (models.Category, error) {
	return s.repo.GetCategoryBySource(ctx, source)
}

func (s *DetailsService) GetBrandBySource(ctx context.Context, source string) (models.Brand, error) {
	return s.repo.GetBrandBySource(ctx, source)
}

func (s *DetailsService) GetModelBySource(ctx context.Context, source string) (models.Model, error) {
	return s.repo.GetModelBySource(ctx, source)
}

func (s *DetailsService) GetColorBySource(ctx context.Context, source string) (models.Color, error) {
	return s.repo.GetColorBySource(ctx, source)
}

func (s *DetailsService) GetBodyBySource(ctx context.Context, source string) (models.Body, error) {
	return s.repo.GetBodyBySource(ctx, source)
}

func (s *DetailsService) GetGenerationBySource(ctx context.Context, source string) (models.Generation, error) {
	return s.repo.GetGenerationBySource(ctx, source)
}

func (s *DetailsService) GetCityBySource(ctx context.Context, source string) (models.City, error) {
	return s.repo.GetCityBySource(ctx, source)
}

func (s *DetailsService) AddSourceUrl(ctx context.Context, photo models.CarPhoto) error {
	return s.repo.AddSourceUrl(ctx, photo)
}
