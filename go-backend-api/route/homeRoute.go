package route

import (
	"github.com/SoumyadipBhowmik/go-backend/controller"
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(app *fiber.App) {
	api := app.Group("/home")
	api.Get("", controller.ShowHomePage)
}
