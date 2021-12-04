package inmemory

import (
	"api/internal/models"
	"context"
	"fmt"
	"sync"
)

type UserRepo struct {
	data map[int]*models.User

	mu *sync.RWMutex
}

func (db *UserRepo) Create(ctx context.Context, user *models.User) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[user.ID] = user
	return nil
}

func (db *UserRepo) All(ctx context.Context) ([]*models.User, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	users := make([]*models.User, 0, len(db.data))
	for _, user := range db.data {
		fmt.Printf("%+v", models.Role(user.Role))
		users = append(users, user)
	}

	return users, nil
}

func (db *UserRepo) ByID(ctx context.Context, id int) (*models.User, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	user, ok := db.data[id]
	if !ok {
		return nil, fmt.Errorf("no user with id %d", id)
	}

	return user, nil
}

func (db *UserRepo) Update(ctx context.Context, user *models.User) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[user.ID] = user
	return nil
}

func (db *UserRepo) Delete(ctx context.Context, id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.data, id)
	return nil
}
