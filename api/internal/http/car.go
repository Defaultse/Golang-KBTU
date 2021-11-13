package http

import (
	"api/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func (s *Server) createCar(w http.ResponseWriter, r *http.Request) {
	car := new(models.Car)
	if err := json.NewDecoder(r.Body).Decode(car); err != nil {
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	s.store.Cars().Create(r.Context(), car)
	fmt.Println("Created", *car)
}

func (s *Server) getAllCars(w http.ResponseWriter, r *http.Request) {
	cars, err := s.store.Cars().All(r.Context())
	if err != nil {
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	render.JSON(w, r, cars)
}

func (s *Server) getCarById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	car, err := s.store.Cars().ByID(r.Context(), id)
	if err != nil {
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	render.JSON(w, r, car)
}

func (s *Server) updateCar(w http.ResponseWriter, r *http.Request) {
	car := new(models.Car)
	if err := json.NewDecoder(r.Body).Decode(car); err != nil {
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	s.store.Cars().Update(r.Context(), car)
}

func (s *Server) deleteCarByid(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	s.store.Cars().Delete(r.Context(), id)
}
