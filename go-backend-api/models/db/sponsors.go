package db

import "github.com/google/uuid"

type Sponsors struct {
	ID           uuid.UUID `db:"id"`
	Organisation *string   `db:"organisation"`
	Individual   *string   `db:"individual"`
}
