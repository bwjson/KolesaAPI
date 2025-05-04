package service

import (
	"context"
	"github.com/bwjson/kolesa_api/internal/models"
	"github.com/bwjson/kolesa_api/internal/repository"
)

type UsersService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *UsersService {
	return &UsersService{repo: repo}
}

func (s *UsersService) Create(ctx context.Context, user models.User) (int, error) {
	return s.repo.Create(ctx, user)
}

func (s *UsersService) GetAll(ctx context.Context) ([]models.User, error) {
	return s.repo.GetAll(ctx)
}

func (s *UsersService) GetByID(ctx context.Context, id int) (models.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UsersService) GetByPhoneNumber(ctx context.Context, phoneNumber string) (*models.User, error) {
	return s.repo.GetByPhoneNumber(ctx, phoneNumber)
}

func (s *UsersService) Update(ctx context.Context, id int, user models.User) error {
	return s.repo.Update(ctx, id, user)
}

func (s *UsersService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
