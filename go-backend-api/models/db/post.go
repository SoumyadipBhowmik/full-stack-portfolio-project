package db

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	Id          uuid.UUID  `db:"id"`
	UserId      User       `db:"user_id"`
	ImageUrl    string     `db:"image_url"`
	Description string     `db:"description"`
	Reactions   int        `db:"reactions"`
	Comments    *[]Comment `db:"comments"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
}
