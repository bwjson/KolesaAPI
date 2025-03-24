package repository

import (
	"context"
	"errors"
	"github.com/bwjson/api/internal/dto"
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

func (r *CarsRepo) GetAllExtended(ctx context.Context, limit, offset int) ([]dto.Car, int, error) {
	var cars []dto.Car

	res := r.db.WithContext(ctx).
		Preload("User").
		Preload("Category").
		Preload("Brand").
		Preload("Color").
		Preload("Generation").
		Preload("Body").
		Preload("City").
		Limit(limit).
		Offset(offset).
		Find(&cars)

	if res.Error != nil {
		return nil, 0, errors.New("No cars found")
	}

	return cars, int(res.RowsAffected), nil
}

func (r *CarsRepo) GetAll(ctx context.Context, limit, offset int) ([]dto.Car, int, error) {
	var cars []dto.Car

	res := r.db.WithContext(ctx).
		Select("cars.id", "cars.price", "cars.category_id", "cars.brand_id", "cars.model_id").
		Preload("Category").
		Preload("Brand").
		Preload("Model").
		Limit(limit).
		Offset(offset).
		Find(&cars)

	if res.Error != nil {
		return nil, 0, errors.New("No cars found")
	}

	return cars, int(res.RowsAffected), nil
}

func (r *CarsRepo) GetById(ctx context.Context, id int) (dto.Car, error) {
	//var good dto.Good
	//
	//err := r.db.QueryRow("SELECT id, name, description, photo_url, price FROM goods WHERE id = $1", id).
	//	Scan(&good.Id, &good.Name, &good.Description, &good.PhotoUrl, &good.Price)
	//
	//if err != nil {
	//	return good, err
	//}
	//
	return dto.Car{}, nil
}

func (r *CarsRepo) UpdateById(ctx context.Context, id int, good dto.Car) error {
	return nil
}

func (r *CarsRepo) DeleteById(ctx context.Context, id int) error {
	return nil
}
