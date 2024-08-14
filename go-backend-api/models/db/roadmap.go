package db

import (
	"time"

	"github.com/google/uuid"
)

type RoadMap struct {
	ID           uuid.UUID `db:"id"`
	Name         string    `db:"name"`
	Start        time.Time `db:"start_time"`
	ExpectedTime time.Time `db:"expected_end_date"`
	Active       bool      `db:"active"`
	LeadBy       string    `db:"lead_by"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
