package postgres

import (
	"api/internal/models"
	"api/internal/store"
	"context"
	"github.com/jmoiron/sqlx"
)

func (db *DB) Cars() store.CarsRepository {
	if db.cars == nil {
		db.cars = NewCarsRepository(db.conn)
	}

	return db.cars
}

type CarsRepository struct {
	conn *sqlx.DB
}

func NewCarsRepository(conn *sqlx.DB) store.CarsRepository {
	return &CarsRepository{conn: conn}
}

// CRUD

func (c CarsRepository) Create(ctx context.Context, car *models.Car) error {
	_, err := c.conn.Exec("INSERT INTO cars VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", car.OwnerId, car.Type, car.Brand, car.Model, car.Year, car.EngineCapacity, car.Description, car.Price, car.Available)
	if err != nil {
		return err
	}

	return nil
}

func (c CarsRepository) All(ctx context.Context) ([]*models.Car, error) {
	cars := make([]*models.Car, 0)
	if err := c.conn.Select(&cars, "SELECT * FROM cars"); err != nil {
		return nil, err
	}

	return cars, nil
}

func (c CarsRepository) ByID(ctx context.Context, id int) (*models.Car, error) {
	car := new(models.Car)
	if err := c.conn.Get(car, "SELECT * FROM cars WHERE id=$1", id); err != nil {
		return nil, err
	}

	return car, nil
}

func (c CarsRepository) Update(ctx context.Context, car *models.Car) error {
	_, err := c.conn.Exec("UPDATE cars SET price = $1, available = $2 WHERE id=$3", car.Price, car.Available, car.ID)
	if err != nil {
		return err
	}

	return nil
}

func (c CarsRepository) Delete(ctx context.Context, id int) error {
	_, err := c.conn.Exec("DELETE FROM cars WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}


