package main

import (
	"context"
	"user-api/internal/http"
	"user-api/internal/store/postgres"
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
		http.WithAddress(":8001"),
		http.WithStore(store),
	//http.WithElastic(res),
	)
	if err := srv.Run(); err != nil {
		panic(err)
	}

	srv.WaitForGracefulTermination()

}
