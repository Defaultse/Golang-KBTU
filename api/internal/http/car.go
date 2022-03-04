package http

import (
	"api/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/cache/v8"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func (s *Server) createCar(w http.ResponseWriter, r *http.Request) {
	car := new(models.Car)
	if err := json.NewDecoder(r.Body).Decode(car); err != nil {
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}
	car.OwnerId = VerifyToken(w, r)
	if err := s.store.Cars().Create(r.Context(), car); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}
	fmt.Println("Created", *car)
}

func (s *Server) getTopList(w http.ResponseWriter, r *http.Request) {
	var top []*models.Car
	ctx := context.TODO()
	key := "TopCars"

	if err := s.mycache.Get(ctx, key, &top); err == nil {
		render.JSON(w, r, top)
	}

	var cached_cars []*models.Car
	if err := s.mycache.Get(s.ctx, "TopCars", &cached_cars); err == nil {
		fmt.Println(cached_cars)
	}
}

func (s *Server) getAllCars(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("description") != "" {
		// ElasticSearch for description
		idStr := chi.URLParam(r, "description")
		description1 := r.URL.Query().Get("description")
		print(r.Form.Get("description"))
		print(description1)
		car, err := s.store.Cars().Search(r.Context(), &idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		render.JSON(w, r, car)
	}
	cars, err := s.store.Cars().All(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}

	render.JSON(w, r, cars)
}

func (s *Server) listByBrand(w http.ResponseWriter, r *http.Request) {
	brand := chi.URLParam(r, "brand")
    cars, err := s.store.Cars().ByBrand(r.Context(), brand)
	if err != nil {
		fmt.Fprintf(w, "Unknown err[http]: %v", err)
		return
	}

	if err := s.mycache.Set(&cache.Item{
		Ctx: s.ctx,
		Key: "TopCars",
		Value: cars,
		TTL:   time.Minute,
	}); err != nil{
		panic(err)
	}

	render.JSON(w, r, cars)
}

func (s *Server) listByBrandAndModel(w http.ResponseWriter, r *http.Request) {
	brand := chi.URLParam(r, "brand")
	model := chi.URLParam(r, "model")
	cars, err := s.store.Cars().ByBrandAndModel(r.Context(), brand, model)
	if err != nil {
		fmt.Fprintf(w, "Unknown err[http]: %v", err)
		return
	}

	if err := s.mycache.Set(&cache.Item{
		Ctx: s.ctx,
		Key: "TopCars",
		Value: cars,
		TTL:   time.Minute,
	}); err != nil{
		panic(err)
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
	ownerId := VerifyToken(w, r)
	if err := json.NewDecoder(r.Body).Decode(car); err != nil {
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	if car.OwnerId == ownerId {
		s.store.Cars().Update(r.Context(), car)
	}
}

func (s *Server) deleteCarByid(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	//Fetching current cached data
	var cached_cars []*models.Car
	if err := s.mycache.Get(s.ctx, "TopCars", &cached_cars); err == nil {
		fmt.Println(cached_cars)
	}

	//Deleting from DB
	err = s.store.Cars().Delete(r.Context(), id)
	if err != nil {
		return
	}

	//Deleting from cache
	for i, v := range cached_cars {
		if id == v.ID {
			cached_cars = append(cached_cars[:i], cached_cars[i+1:]...)
		}
		if err := s.mycache.Set(&cache.Item{
			Ctx: s.ctx,
			Key: "TopCars",
			Value: cached_cars,
			TTL:   time.Minute,
		}); err != nil{
			panic(err)
		}
	}
}
