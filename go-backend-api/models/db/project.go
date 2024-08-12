package db

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID          uuid.UUID   `db:"id"`
	Name        string      `db:"name"`
	CoverUrl    string      `db:"description"`
	Images      *[]string   `db:"images"`
	Active      bool        `db:"active"`
	Topic       string      `db:"topic"`
	RoadMap     *RoadMap    `db:"road_map"`
	Volunteers  *[]User     `db:"volunteers"`
	Sponsors    *[]Sponsors `db:"sponsors"`
	Description *string     `db:"description"`
	StartTIme   time.Time   `db:"start_time"`
	EndTime     time.Time   `db:"end_time"`
	CreatedAt   time.Time   `db:"created_at"`
	UpdatedAt   time.Time   `db:"updated_at"`
}
