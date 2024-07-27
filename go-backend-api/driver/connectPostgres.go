package driver

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectToPostgresDB(ctx context.Context) *pgxpool.Pool {
	conn, err := pgxpool.New(ctx, os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}
	return conn
}
