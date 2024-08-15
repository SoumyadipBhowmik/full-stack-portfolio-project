package db

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	Id        uuid.UUID `db:"id"`
	UserId    User      `db:"user_id"`
	ImageUrl  string    `db:"image_url"`
	Text      string    `db:"text"`
	Reactions int       `db:"reactions"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
