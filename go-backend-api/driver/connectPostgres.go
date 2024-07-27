package driver

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectToPostgresDB(ctx context.Context) *pgxpool.Pool {
	conn, err := pgxpool.New(ctx, os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}
	ping := conn.Ping(ctx)
	if ping != nil {
		fmt.Print("couldn't establish connection to database ", ping)
	}
	return conn
}
