package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user-api/internal/models"
)

func (s *Server) createOwnerCars(w http.ResponseWriter, r *http.Request) {
	ownerCars := new(models.OwnerCars)
	if err := json.NewDecoder(r.Body).Decode(ownerCars); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	if err := s.store.OwnerCars().Create(r.Context(), ownerCars); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}
}
