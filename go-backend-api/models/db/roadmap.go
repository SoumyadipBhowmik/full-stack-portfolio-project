package db

import (
	"time"

	"github.com/google/uuid"
)

type RoadMap struct {
	ID           uuid.UUID `db:"id"`
	Name         string    `db:"name"`
	Start        string    `db:"start"`
	ExpectedTime string    `db:"expected_time"`
	Active       bool      `db:"active"`
	LeadBy       string    `db:"lead_by"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
