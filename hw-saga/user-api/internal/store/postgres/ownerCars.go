package postgres

import (
"context"
"github.com/jmoiron/sqlx"
"user-api/internal/models"
"user-api/internal/store"
)

func (db *DB) OwnerCars() store.OwnerCarsRepository {
	if db.ownerCars == nil {
		db.ownerCars = NewOwnerCarsRepository(db.conn)
	}

	return db.ownerCars
}

type OwnerCarsRepository struct {
	conn *sqlx.DB
}

func NewOwnerCarsRepository(conn *sqlx.DB) store.OwnerCarsRepository {
	return &OwnerCarsRepository{conn: conn}
}

// CRUD

func (u OwnerCarsRepository) Create(ctx context.Context, ownerCars *models.OwnerCars) error {
	_, err := u.conn.Exec("INSERT INTO owner_cars(owner_id, car_id) VALUES ($1, $2)", ownerCars.UserID, ownerCars.CarId)
	if err != nil {
		return err
	}

	return nil
}
