package main

import (
	"api/internal/http"
	"api/internal/store/postgres"
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/cache/v8"
	"log"
	"time"
)

func main() {
	url := "postgres://postgres:admin@localhost:5432/mytulpar 1.0"
	store := postgres.NewDB()
	if err := store.Connect(url); err != nil {
		panic(err)
	}
	defer store.Close()

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	mycache := cache.New(&cache.Options{
		LocalCache: cache.NewTinyLFU(100, time.Minute),
	})

	srv := http.NewServer(
		context.Background(),
		http.WithAddress(":8000"),
		http.WithStore(store),
		http.WithRedisCache(mycache),
	    http.WithElastic(es),
	)
	if err := srv.Run(); err != nil {
		panic(err)
	}

	srv.WaitForGracefulTermination()
}
