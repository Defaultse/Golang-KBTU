package postgres

import (
	"bytes"
	"car-api/internal/models"
	"car-api/internal/store"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	esv8 "github.com/elastic/go-elasticsearch/v8"
	esv8api "github.com/elastic/go-elasticsearch/v8/esapi"
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

// elastic search
func (c CarsRepository) Index(ctx context.Context, car *models.Car) error {
	body := models.Car{
		ID:          car.ID,
		Description: car.Description,
	}

	var buf bytes.Buffer

	_ = json.NewEncoder(&buf).Encode(body) // XXX: error omitted

	req := esv8api.IndexRequest{
		Index:      c.index,
		Body:       &buf,
		DocumentID: car.Description,
		Refresh:    "true",
	}

	resp, err := req.Do(ctx, c.client) // XXX: error omitted
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	io.Copy(ioutil.Discard, resp.Body)

	return nil
}

func (c CarsRepository) Search(ctx context.Context, description *string) ([]models.Car, error) {
	if description == nil {
		return nil, nil
	}

	should := make([]interface{}, 0, 3)

	if description != nil {
		should = append(should, map[string]interface{}{
			"match": map[string]interface{}{
				"description": *description,
			},
		})
	}

	var query map[string]interface{}

	if len(should) > 1 {
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"should": should,
				},
			},
		}
	} else {
		query = map[string]interface{}{
			"query": should[0],
		}
	}

	var buf bytes.Buffer

	_ = json.NewEncoder(&buf).Encode(query)

	req := esv8api.SearchRequest{
		Index: []string{c.index},
		Body:  &buf,
	}

	resp, _ := req.Do(ctx, c.client) // XXX: error omitted
	defer resp.Body.Close()

	var hits struct {
		Hits struct {
			Hits []struct {
				Source models.Car `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	_ = json.NewDecoder(resp.Body).Decode(&hits) // XXX: error omitted

	res := make([]models.Car, len(hits.Hits.Hits))

	for i, hit := range hits.Hits.Hits {
		res[i].ID = hit.Source.ID
		res[i].Description = hit.Source.Description
	}

	return res, nil
}

// CRUD
func (c CarsRepository) Create(ctx context.Context, car *models.Car) error {
	_, err := c.conn.Exec("INSERT INTO cars VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", car.OwnerId, car.RentTypeId, car.CarTypeId, car.ModelId, car.ColorId, car.Year, car.EngineCapacity, car.Description, car.Price)
	if err != nil {
		return err
	}
	c.Index(ctx, car)
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

func (c CarsRepository) ByID(ctx context.Context, id int) (*models.Car, error) {
	car := new(models.Car)
	if err := c.conn.Get(car, "SELECT * FROM cars WHERE id=$1", id); err != nil {
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
