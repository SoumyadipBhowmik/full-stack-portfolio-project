package controllers

import (
	"log"

	"github.com/SoumyadipBhowmik/go-backend/models/dto"
	"github.com/SoumyadipBhowmik/go-backend/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PostController struct {
	service *services.PostServices
}

func NewPostController(service *services.PostServices) *PostController {
	return &PostController{service: service}
}

func (controller *PostController) CreatePost(context *fiber.Ctx) error {
	var post dto.PostDTO
	err := context.BodyParser(&post)
	if err != nil {
		log.Fatal(err.Error())
	}
	return context.JSON(controller.service.CreatePost(post.UserId, post.ImageUrl, post.Description))

}

func (controller *PostController) FetchPost(context *fiber.Ctx) error {
	queryId := context.Params("id")
	passId, _ := uuid.Parse(queryId)
	return context.JSON(controller.service.FetchPost(passId))
}

func (controller *PostController) FetchAll(context *fiber.Ctx) error {
	return context.JSON(controller.service.FetchAll())
}

func (controller *PostController) FetchUserPosts(context *fiber.Ctx) error {
	var userId uuid.UUID
	err := context.BodyParser(&userId)
	if err != nil {
		log.Fatalf("can't use body parser because: %v", err.Error())
	}
	return context.JSON(controller.service.FetchUserPosts(userId))
}
