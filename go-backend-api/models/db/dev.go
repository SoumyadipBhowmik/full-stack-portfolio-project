package db

import (
	"time"

	"github.com/google/uuid"
)

type Dev struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Role        string    `db:"role"`
	Description string    `db:"description"`
	Github      string    `db:"github"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"created_at"`
}
