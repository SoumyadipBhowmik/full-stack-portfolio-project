package models

import "github.com/google/uuid"

type Dev struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Role        string    `db:"role"`
	Description string    `db:"description"`
	Github      string    `db:"github"`
}
