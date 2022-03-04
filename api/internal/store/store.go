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

type AuthRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	generatePasswordHash(password string) string
}

type CarsRepository interface {
	Create(ctx context.Context, car *models.Car) error
	All(ctx context.Context) ([]*models.FullCar, error)
	ByBrand(ctx context.Context, brand string) ([]*models.Car, error)
	ByBrandAndModel(ctx context.Context, brand string, model string) ([]*models.Car, error)
	ByID(ctx context.Context, id int) (*models.FullCar, error)
	Update(ctx context.Context, car *models.Car) error
	Delete(ctx context.Context, id int) error

	Index(ctx context.Context, car *models.FullCar) error
	Search(ctx context.Context, description *string) ([]*models.FullCar, error)
}

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	All(ctx context.Context) ([]*models.User, error)
	GetUser(ctx context.Context, email string, password_hash string) (*models.User, error)
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
