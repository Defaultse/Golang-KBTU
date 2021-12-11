package saga

import (
	"car-api/internal/models"
	"encoding/json"
	"log"
	"net/http"
)

type (
	CompensationRequest struct {
		CarID int `json:"car_id"`
	}

	Response struct {
		CarID int  `json:"car_id,omitempty"`
		Success bool `json:"success"`
	}
)

func creationSuccess(w http.ResponseWriter, r *http.Request) {
	var payload models.Car
	json.NewDecoder(r.Body).Decode(&payload)

	log.Printf("[car ID %s] creation: success\n", payload.ID)

	resp := Response{
		CarID: payload.ID,
		Success: true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func creationFailed(w http.ResponseWriter, r *http.Request) {
	var payload models.Car
	json.NewDecoder(r.Body).Decode(&payload)

	log.Printf("Creation of car %s : FAILED!!!\n", payload.ID)

	resp := Response{
		Success: false,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(resp)
}

func creationCompensation(w http.ResponseWriter, r *http.Request) {
	var payload CompensationRequest
	json.NewDecoder(r.Body).Decode(&payload)

	log.Printf("[rollback] rollback car_id %d : success\n", payload.CarID)

	resp := Response{
		Success: true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}