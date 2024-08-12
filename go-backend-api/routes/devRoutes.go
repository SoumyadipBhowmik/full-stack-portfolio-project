package routes

import (
	"github.com/SoumyadipBhowmik/go-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func initializeDevRoutes(app *fiber.App, devController *controllers.ContributerController) {
	api := app.Group("/dev")
	api.Post("", devController.AddContributer)
}
