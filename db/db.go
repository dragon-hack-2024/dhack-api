package db

import (
	"context"
	"dhack-api/model"
	"log"

	"github.com/jackc/pgx/v5"
)

type Store struct {
	conn    *pgx.Conn
	Queries *model.Queries
}

func Connect(source string) (*Store, error) {
	conn, err := pgx.Connect(context.Background(), source)
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

func (store *Store) Close() error {
	return store.conn.Close(context.Background())
}
