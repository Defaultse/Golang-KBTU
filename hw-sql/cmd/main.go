package main

import (
	"api/internal/http"
	"api/internal/store/postgres"
	"context"
)

func main() {
	//ctx := context.Background()
	url := "postgres://postgres:admin@localhost:5432/mytulpar"
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

	//elasticsearch connection
	//es, err := elasticsearch.NewDefaultClient()
	//if err != nil {
	//	log.Fatalf("Error creating the client: %s", err)
	//}
	//log.Println(elasticsearch.Version)

	//res, err := es.Info()
	//if err != nil {
	//	log.Fatalf("Error getting response: %s", err)
	//}
	//defer res.Body.Close()


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
