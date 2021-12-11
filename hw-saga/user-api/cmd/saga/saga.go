package saga

import (
	"encoding/json"
	"log"
	"net/http"
)

type (

	Request struct {
		UserID      int    `json:"user_id" db:"user_id"`
		CarId      int    `json:"car_id" db:"car_id"`
		Car struct {
			ID      int    `json:"id" db:"id"`
			OwnerId int    `json:"owner_id" db:"owner_id"`
			RentTypeId    int   `json:"rent_type_id" db:"rent_type_id"`
			CarTypeId    int   `json:"car_type_id" db:"car_type_id"`
			ModelId   string `json:"model_id" db:"model_id"`
			ColorId   int    `json:"color_id" db:"color_id"`
			Year           int     `json:"year" db:"year"`
			EngineCapacity float32 `json:"engine_capacity" db:"engine_capacity"`
			Description    string  `json:"description" db:"description"`
			Price          int     `json:"price" db:"price"`
		}
	}

	// CompensationRequest defines compensation request
	CompensationRequest struct {
		UserID      int    `json:"user_id" db:"user_id"`
		CarId      int    `json:"car_id" db:"car_id"`
	}

	// Response defines response
	Response struct {
		UserID      int    `json:"user_id" db:"user_id"`
		CarId      int    `json:"car_id" db:"car_id"`

		Success bool `json:"success"`
	}
)

func creationSuccess(w http.ResponseWriter, r *http.Request) {
	var payload Request
	json.NewDecoder(r.Body).Decode(&payload)

	log.Printf("[User ID %d] added %d : success\n", payload.UserID, payload.CarId)

	resp := Response{
		UserID: payload.UserID,
		Success: true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}


func creationFailed(w http.ResponseWriter, r *http.Request) {
	var payload Request
	json.NewDecoder(r.Body).Decode(&payload)

	log.Printf("[User ID %d] added %d : fail!!!\n", payload.UserID, payload.CarId)

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

	log.Printf("[User ID %d] added %d : rollback\n", payload.UserID, payload.CarId)

	resp := Response{
		Success: true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}