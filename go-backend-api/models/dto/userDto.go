package models

import (
	"github.com/google/uuid"
)

type UserDTO struct {
	ID       uuid.UUIDs `json:"id"`
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Feedback string     `json:"feedback"`
}
