package postgres

import (
	"api/internal/models"
	"api/internal/store"
	"context"
	"fmt"
	esv8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/jmoiron/sqlx"
)

func (db *DB) Cars() store.CarsRepository {
	if db.cars == nil {
		db.cars = NewCarsRepository(db.conn, nil)
	}

	return db.cars
}

type CarsRepository struct {
	conn   *sqlx.DB
	client *esv8.Client
	index  string
}

func NewCarsRepository(conn *sqlx.DB, client *esv8.Client) store.CarsRepository {
	return &CarsRepository{
		conn:   conn,
		client: client,
		index:  "description",
	}
}

// CRUD
func (c CarsRepository) Create(ctx context.Context, car *models.Car) error {
	_, err := c.conn.Exec("INSERT INTO cars (owner_id, rent_type_id, car_type_id, model_id, color_id, year, engine_capacity, description, price) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		car.OwnerId,
		car.RentTypeId,
		car.CarTypeId,
		car.ModelId,
		car.ColorId,
		car.Year,
		car.EngineCapacity,
		car.Description,
		car.Price)
	if err != nil {
		return err
	}

	fmt.Println(car)
	return nil
}

func (c CarsRepository) All(ctx context.Context) ([]*models.FullCar, error) {
	cars := make([]*models.FullCar, 0)
	if err := c.conn.Select(&cars, "SELECT * FROM get_full_cars"); err != nil {
		return nil, err
	}

	return cars, nil
}

func (c CarsRepository) ByBrand(ctx context.Context, brand string) ([]*models.Car, error) {
	cars := make([]*models.Car, 0)
	if err := c.conn.Select(&cars, "Select cars.* from cars join brands on (cars.brand_id=brands.id AND brands.brand=$1)", brand); err != nil {
		return nil, err
	}

	return cars, nil
}

func (c CarsRepository) ByBrandAndModel(ctx context.Context, brand string, model string) ([]*models.Car, error) {
	cars := make([]*models.Car, 0)
	if err := c.conn.Select(&cars, "Select cars.* from cars join models on cars.model_id=models.id AND models.model=$1", model); err != nil {
		return nil, err
	}

	return cars, nil
}

func (c CarsRepository) ByID(ctx context.Context, id int) (*models.FullCar, error) {
	car := new(models.FullCar)
	if err := c.conn.Get(car, "SELECT * FROM get_full_cars WHERE id=$1", id); err != nil {
		return nil, err
	}

	return car, nil
}

func (c CarsRepository) Update(ctx context.Context, car *models.Car) error {
	_, err := c.conn.Exec("UPDATE cars SET price = $1 WHERE id=$3", car.Price, car.ID)
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
