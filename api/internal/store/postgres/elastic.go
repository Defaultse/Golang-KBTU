package postgres

import (
	"api/internal/models"
	"bytes"
	"context"
	"encoding/json"
	esv8api "github.com/elastic/go-elasticsearch/v8/esapi"
	"io"
	"io/ioutil"
)

// elastic search
func (c CarsRepository) Index(ctx context.Context, car *models.FullCar) error {
	body := models.FullCar{
		ID:          car.ID,
		Description: car.Description,
	}

	var buf bytes.Buffer

	_ = json.NewEncoder(&buf).Encode(body)

	req := esv8api.IndexRequest{
		Index:      c.index,
		Body:       &buf,
		DocumentID: car.Description,
		Refresh:    "true",
	}

	resp, err := req.Do(ctx, c.client)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	io.Copy(ioutil.Discard, resp.Body)

	return nil
}

func (c CarsRepository) Search(ctx context.Context, description *string) ([]*models.FullCar, error) {
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

	resp, _ := req.Do(ctx, c.client)
	defer resp.Body.Close()

	var hits struct {
		Hits struct {
			Hits []struct {
				Source models.Car `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	_ = json.NewDecoder(resp.Body).Decode(&hits)

	res := make([]*models.FullCar, len(hits.Hits.Hits))

	for i, hit := range hits.Hits.Hits {
		res[i].ID = hit.Source.ID
		res[i].Description = hit.Source.Description
	}

	return res, nil
}
