package database

import (
	"context"
	"fmt"
	"pesthub/entities"
	"pesthub/features/pest"

	"github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxpool.Pool

const uri = "postgres://greenhubusr:greenhubpw@localhost:5432/greenhubdb"

func init() {
	if connection, err := pgxpool.Connect(context.Background(), uri); err == nil {
		db = connection
	} else {
		message := fmt.Errorf("failed to connect to database. Reason: %v", err)
		panic(message)
	}
}

var Insert pest.InsertCommand = func(pest *entities.Pest) (int64, error) {
	ctx := context.Background()
	query := "INSERT INTO praga (name_common) VALUES ($1) RETURNING id"
	row := db.QueryRow(
		ctx,
		query,
		pest.CommonName,
	)

	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
