package repository

import (
	"context"
	"github.com/bwjson/kolesa_api/internal/dto"
	"gorm.io/gorm"
)

type UsersRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UsersRepo {
	return &UsersRepo{db: db}
}

func (r *UsersRepo) Create(ctx context.Context, user dto.User) error {
	return r.db.WithContext(ctx).Create(&user).Error
}

func (r *UsersRepo) GetAll(ctx context.Context) ([]dto.User, error) {
	var users []dto.User
	err := r.db.WithContext(ctx).Find(&users).Error
	return users, err
}

func (r *UsersRepo) GetByID(ctx context.Context, id int) (dto.User, error) {
	var user dto.User
	err := r.db.WithContext(ctx).First(&user, id).Error
	return user, err
}

func (r *UsersRepo) Update(ctx context.Context, user dto.User) error {
	return r.db.WithContext(ctx).Save(&user).Error
}

func (r *UsersRepo) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&dto.User{}, id).Error
}
