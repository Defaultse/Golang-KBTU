package main

import (
	"car-api/internal/http"
	"car-api/internal/store/postgres"
	"context"
)

func main() {
	//ctx := context.Background()
	url := "postgres://postgres:admin@localhost:5432/mytulpar 1.0"
	//conn, err := pgx.Connect(ctx, urlExample)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	//	os.Exit(1)
	//}
	//defer conn.Close(context.Background())
	//
	//if err := conn.Ping(ctx); err != nil {
	//	panic(err)
	//}
	store := postgres.NewDB()
	if err := store.Connect(url); err != nil {
		panic(err)
	}
	defer store.Close()

	srv := http.NewServer(
		context.Background(),
		http.WithAddress(":8000"),
		http.WithStore(store),
	//http.WithElastic(res),
	)
	if err := srv.Run(); err != nil {
		panic(err)
	}

	srv.WaitForGracefulTermination()
}
