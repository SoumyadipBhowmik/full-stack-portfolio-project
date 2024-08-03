package models

import "github.com/google/uuid"

type Service struct {
	ID          uuid.UUID `db:"uuid"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Icon        string    `db:"icon"`
	Priority    int32     `db:"priority"`
}
