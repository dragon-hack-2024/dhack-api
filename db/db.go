package db

import (
	"context"
	"dhack-api/model"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	conn    *pgxpool.Pool
	Queries *model.Queries
}

func Connect(source string) (*Store, error) {
	conn, err := pgxpool.New(context.Background(), source)
	if err != nil {
		return nil, err
	}

	store := &Store{
		conn:    conn,
		Queries: model.New(conn),
	}

	log.Println("Connected to database")
	return store, err
}

func (store *Store) PingDB() error {
	return store.conn.Ping(context.Background())
}

func (store *Store) Close() {
	store.conn.Close()
}
