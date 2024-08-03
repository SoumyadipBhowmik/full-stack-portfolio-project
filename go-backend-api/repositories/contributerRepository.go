package repositories

import "github.com/jackc/pgx/v5/pgxpool"

type ContributerRepostory struct {
	DB *pgxpool.Pool
}
