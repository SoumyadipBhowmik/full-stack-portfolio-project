package repositories

import (
	"context"
	"log"

	"github.com/SoumyadipBhowmik/go-backend/models/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ContributerRepostory struct {
	DB *pgxpool.Pool
}

func NewContributerRepository(db *pgxpool.Pool) *ContributerRepostory {
	return &ContributerRepostory{DB: db}
}

func (repo *ContributerRepostory) AddContributor(name, role, description, github string) *db.Dev {
	query := `
	INSERT INTO dev (name, role, description, github)
	VALUES ($1, $2, $3, $4)
	RETURNING name, role, description, github;
	`
	var dev db.Dev
	err := repo.DB.QueryRow(context.Background(), query, name, role, description, github).Scan(&dev.Name, &dev.Role, &dev.Description, &dev.Github)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return &dev
}
