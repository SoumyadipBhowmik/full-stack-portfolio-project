package dto

type PostDTO struct {
	UserId      UserDTO `json:"user_id"`
	ImageUrl    string  `json:"image_url"`
	Description string  `json:"description"`
}
