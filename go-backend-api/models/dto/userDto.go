package dto

import "github.com/google/uuid"

type UserDTO struct {
	Id       *uuid.UUID `json:"uuid"`
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Feedback string     `json:"feedback"`
}
