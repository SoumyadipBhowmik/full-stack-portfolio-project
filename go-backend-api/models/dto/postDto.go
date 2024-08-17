package dto

import (
	"time"

	"github.com/google/uuid"
)

type PostDTO struct {
	Id          uuid.UUID `json:"id"`
	UserId      uuid.UUID `json:"user_id"`
	ImageUrl    string    `json:"image_url"`
	Description string    `json:"description"`
	Reactions   *int      `json:"reactions"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
