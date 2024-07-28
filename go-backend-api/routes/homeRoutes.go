package routes

import (
	"github.com/SoumyadipBhowmik/go-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(app *fiber.App) {
	api := app.Group("/home")
	api.Get("", controllers.ShowHomeCards)
	api.Post("", controllers.CreateCards)
}
