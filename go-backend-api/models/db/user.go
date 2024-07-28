package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUIDs `db:"id"`
	Name      string     `db:"name"`
	Email     string     `db:"email"`
	Feedback  string     `db:"feedback"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
}
