package routes

import (
	"github.com/SoumyadipBhowmik/go-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func initializePostRoutes(app *fiber.App, postController *controllers.PostController) {

	api := app.Group("/posts")
	api.Post("", postController.CreatePost)
	api.Get("", postController.FetchAll)
	api.Get(":id", postController.FetchPost)
	api.Get("users/:id", postController.FetchUserPosts)
}
