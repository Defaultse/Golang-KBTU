package inmemory

import (
	"api/internal/models"
	"api/internal/store"
	"sync"
)

type DB struct {
	carsRepo     store.CarsRepository
	userRepo     store.UserRepository
	feedbackRepo store.FeedbackRepository

	mu *sync.RWMutex
}

func NewDB() store.Store {
	return &DB{
		mu: new(sync.RWMutex),
	}
}

func (db *DB) Cars() store.CarsRepository {
	if db.carsRepo == nil {
		db.carsRepo = &CarsRepo{
			data: make(map[int]*models.Car),
			mu:   new(sync.RWMutex),
		}
	}

	return db.carsRepo
}

func (db *DB) User() store.UserRepository {
	if db.userRepo == nil {
		db.userRepo = &UserRepo{
			data: make(map[int]*models.User),
			mu:   new(sync.RWMutex),
		}
	}

	return db.userRepo
}

func (db *DB) Feedback() store.FeedbackRepository {
	if db.feedbackRepo == nil {
		db.feedbackRepo = &FeedbackRepo{
			data: make(map[int]*models.Feedback),
			mu:   new(sync.RWMutex),
		}
	}

	return db.feedbackRepo
}
