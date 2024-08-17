package services

import (
	"github.com/SoumyadipBhowmik/go-backend/models/dto"
	"github.com/SoumyadipBhowmik/go-backend/repositories"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type PostServices struct {
	repo *repositories.PostRepository
}

func NewPostServices(repo *repositories.PostRepository) *PostServices {
	return &PostServices{repo: repo}
}

func (service *PostServices) CreatePost(userId uuid.UUID, imageUrl, description string) *dto.PostDTO {

	post := service.repo.CreatePost(userId, imageUrl, description)

	postDto := &dto.PostDTO{
		Id:          post.Id,
		UserId:      post.UserId,
		ImageUrl:    post.ImageUrl,
		Description: post.Description,
		Reactions:   &post.Reactions,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
	}
	return postDto
}

func (service *PostServices) FetchPost(id uuid.UUID) *dto.PostDTO {
	post := service.repo.FetchPost(id)
	var postDto dto.PostDTO
	copier.Copy(&postDto, post)
	return &postDto
}

func (service *PostServices) FetchAll() *[]dto.PostDTO {
	posts := service.repo.FetchAll()
	var postsDto []dto.PostDTO
	copier.Copy(&postsDto, posts)
	return &postsDto
}
