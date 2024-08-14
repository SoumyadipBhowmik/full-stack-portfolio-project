package repositories

import (
	"context"
	"log"
	"time"

	"github.com/SoumyadipBhowmik/go-backend/models/db"
	"github.com/SoumyadipBhowmik/go-backend/models/dto"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jinzhu/copier"
)

type RoadMapRepository struct {
	Db *pgxpool.Pool
}

func NewRoadMapRepository(db *pgxpool.Pool) *RoadMapRepository {
	return &RoadMapRepository{Db: db}
}

func (repo *RoadMapRepository) RoadMapCreation(name string, startDate, expectedEndDate time.Time, active bool, leadBy string) *db.RoadMap {
	query := `
	INSERT INTO road_map (name, start_time, expected_end_date, active, lead_by)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING name, start_time, expected_end_date, active, lead_by;
	`
	var roadmapDTO dto.RoadMapDTO
	err := repo.Db.QueryRow(context.Background(), query, name, startDate, expectedEndDate, active, leadBy).Scan(
		&roadmapDTO.Name,
		&roadmapDTO.Start,
		&roadmapDTO.ExpectedTime,
		&roadmapDTO.Active,
		&roadmapDTO.LeadBy,
	)
	if err != nil {
		log.Fatalf("error in query %v", err.Error())
	}
	var roadmap db.RoadMap
	copyerror := copier.Copy(&roadmap, roadmapDTO)
	if copyerror != nil {
		log.Fatalf("copy error because: %v", copyerror.Error())
	}
	return &roadmap
}
