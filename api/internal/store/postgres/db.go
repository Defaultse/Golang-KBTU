package postgres

import (
	"api/internal/store"
	"fmt"
    "github.com/jmoiron/sqlx"
	"os"
)

type DB struct {
	conn *sqlx.DB

	users store.UserRepository
	cars store.CarsRepository
}

func NewDB() store.Store {
	return &DB{}
}

func (db *DB) Connect(url string) error{
	conn, err := sqlx.Connect("pgx", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	if err := conn.Ping(); err != nil {
		panic(err)
	}

	db.conn = conn
	return nil
}

func (db *DB) Close() error{
	return db.conn.Close()
}

func (D DB) Feedback() store.FeedbackRepository {
	panic("implement me")
}

