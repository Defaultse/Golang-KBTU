package http

import (
	"api/internal/store"
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/cache/v8"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type Server struct {
	ctx         context.Context
	idleConnsCh chan struct{}
	store       store.Store
	es *elasticsearch.Client
	mycache *cache.Cache

	Address string
}

func NewServer(ctx context.Context, opts ...ServerOption) *Server {
	srv := &Server{
		ctx:         ctx,
		idleConnsCh: make(chan struct{}),
	}

	for _, opt := range opts {
		opt(srv)
	}

	return srv
}


func (s *Server) basicHandler() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Route("/cars", func(r chi.Router) {
			r.Get("/", s.getAllCars)
			r.Get("/top", s.getTopList)
			r.Get("/{brand}",s.listByBrand)
			r.Get("/{brand}/{model}", s.listByBrandAndModel)
			r.Get("/show/{id}", s.getCarById)
			r.Post("/", s.createCar)
			r.Put("/", s.updateCar)
			r.Delete("/{id}", s.deleteCarByid)
		})
	})

	//r.Get("/cars?description={d}", s.elasticGetByDescription)
	//r.Post("/cars", s.createCar)
	//r.Get("/cars/{brand}", s.listByBrand)
	//r.Get("/cars/{brand}/{model}", s.listByBrandAndModel)
	//r.Get("/car/show/{id}", s.getCarById)
	//r.Put("/cars", s.updateCar)
	//r.Delete("/car/delete/{id}", s.deleteCarByid)

	r.Post("/user/register", s.registration)
	r.Post("/user/login", s.login)
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
