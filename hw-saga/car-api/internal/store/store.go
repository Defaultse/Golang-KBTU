package store

import (
	"car-api/internal/models"
	"context"
)

type Store interface {
	Connect(url string) error
	Close() error

	Cars() CarsRepository
}

type CarsRepository interface {
	Create(ctx context.Context, car *models.Car) error
	All(ctx context.Context) ([]*models.FullCar, error)
	ByID(ctx context.Context, id int) (*models.Car, error)
	Update(ctx context.Context, car *models.Car) error
	Delete(ctx context.Context, id int) error

	Index(ctx context.Context, car *models.Car) error
	Search(ctx context.Context, description *string) ([]models.Car, error)
}
