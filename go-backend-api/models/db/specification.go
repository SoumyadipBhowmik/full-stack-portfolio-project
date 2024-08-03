package models

import "github.com/google/uuid"

type Specification struct {
	ID       uuid.UUID `db:"id"`
	Type     string    `db:"type"`
	LongText string    `db:"longtext"`
}
