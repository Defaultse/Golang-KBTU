package store

import (
	"context"
	"user-api/internal/models"
)

type Store interface {
	Connect(url string) error
	Close() error

	User() UserRepository
	OwnerCars() OwnerCarsRepository
}

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	All(ctx context.Context) ([]*models.User, error)
	ByID(ctx context.Context, id int) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id int) error
}

type OwnerCarsRepository interface {
	Create(ctx context.Context, user *models.OwnerCars) error
}