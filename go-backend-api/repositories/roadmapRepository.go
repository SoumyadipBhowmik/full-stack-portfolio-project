package repositories

import (
	"github.com/SoumyadipBhowmik/go-backend/models/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProjectRepository struct {
	db *pgxpool.Pool
}

func NewRoadMapRepository(db *pgxpool.Pool) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (repo *ProjectRepository) ProjectCreation(name, start, expectedTime, active, leadBy string) (db.Project, error) {
	query := `
	INSERT INTO project ()

	`
	var newProject db.Project
	return newProject, nil
}
