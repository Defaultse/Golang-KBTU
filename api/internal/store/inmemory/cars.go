package inmemory

import (
	"api/internal/models"
	"context"
	"fmt"
	"sync"
)

type CarsRepo struct {
	data map[int]*models.Car

	mu *sync.RWMutex
}

func (db *CarsRepo) Create(ctx context.Context, car *models.Car) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[car.ID] = car
	return nil
}

func (db *CarsRepo) All(ctx context.Context) ([]*models.Car, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	cars := make([]*models.Car, 0, len(db.data))
	for _, car := range db.data {
		cars = append(cars, car)
	}

	return cars, nil
}

func (db *CarsRepo) ByID(ctx context.Context, id int) (*models.Car, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	car, ok := db.data[id]
	if !ok {
		return nil, fmt.Errorf("no car with id %d", id)
	}

	return car, nil
}

func (db *CarsRepo) Update(ctx context.Context, car *models.Car) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[car.ID] = car
	return nil
}

func (db *CarsRepo) Delete(ctx context.Context, id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.data, id)
	return nil
}
