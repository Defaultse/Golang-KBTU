package store

import (
	"api/internal/models"
	"context"
)

type Store interface {
	Connect(url string) error
	Close() error

	Cars() CarsRepository
	User() UserRepository
	Feedback() FeedbackRepository
}

type CarsRepository interface {
	Create(ctx context.Context, car *models.Car) error
	All(ctx context.Context, filter *models.CarsFilter) ([]*models.Car, error)
	ByID(ctx context.Context, id int) (*models.Car, error)
	Update(ctx context.Context, car *models.Car) error
	Delete(ctx context.Context, id int) error
}

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	All(ctx context.Context) ([]*models.User, error)
	ByID(ctx context.Context, id int) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id int) error
}

type FeedbackRepository interface {
	Create(ctx context.Context, user *models.Feedback) error
	All(ctx context.Context) ([]*models.Feedback, error)
	ByID(ctx context.Context, id int) ([]*models.Feedback, error)
	Update(ctx context.Context, user *models.Feedback) error
	Delete(ctx context.Context, userId int) error
}
