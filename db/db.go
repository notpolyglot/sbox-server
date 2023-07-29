package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Conn *pgxpool.Pool

func Connect() (err error) {
	ctx := context.Background()
	Conn, err = pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return
	}
	return
}
