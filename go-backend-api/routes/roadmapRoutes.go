package routes

import (
	"github.com/SoumyadipBhowmik/go-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func initializeRoadMapRoutes(app *fiber.App, controller *controllers.RoadMapController) {
	api := app.Group("/roadmap")
	api.Post("", controller.RoadMapCreation)
}
