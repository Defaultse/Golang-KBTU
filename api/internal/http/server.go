package http

import (
	"api/internal/store"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type Server struct {
	ctx         context.Context
	idleConnsCh chan struct{}
	store       store.Store

	Address string
}

func NewServer(ctx context.Context, address string, store store.Store) *Server {
	return &Server{
		ctx:         ctx,
		idleConnsCh: make(chan struct{}),
		store:       store,

		Address: address,
	}
}

func (s *Server) basicHandler() chi.Router {
	r := chi.NewRouter()

	r.Post("/cars", s.createCar)
	r.Get("/cars", s.getAllCars)
	r.Get("/cars/{id}", s.getCarById)
	r.Put("/cars", s.updateCar)
	r.Delete("/cars/{id}", s.deleteCarByid)

	r.Post("/user", s.createUser)
	r.Get("/users", s.getAllUsers)
	r.Get("/user/{id}", s.getUserById)
	r.Put("/user", s.updateUser)
	r.Delete("/user/{id}", s.deleteUser)

	r.Post("/feedbacks", s.createFeedback)
	r.Get("/feedbacks/{userId}", s.getFeedbacksbyProfileId)
	r.Put("/feedbacks", s.updateFeedbacks)
	r.Delete("/feedbacks/{id}", s.deleteFeedbackByid)

	return r
}

func (s *Server) Run() error {
	srv := &http.Server{
		Addr:         s.Address,
		Handler:      s.basicHandler(),
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 30,
	}
	go s.ListenCtxForGT(srv)

	log.Println("[HTTP] Server running on", s.Address)
	return srv.ListenAndServe()
}

func (s *Server) ListenCtxForGT(srv *http.Server) {
	<-s.ctx.Done()

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("[HTTP] Got err while shutting down^ %v", err)
	}

	log.Println("[HTTP] Proccessed all idle connections")
	close(s.idleConnsCh)
}

func (s *Server) WaitForGracefulTermination() {
	<-s.idleConnsCh
}
