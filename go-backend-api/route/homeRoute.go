package route

import "github.com/gofiber/fiber/v2"

func InitializeRoutes(app *fiber.App) {
	app.Get("/home", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to routes")
	})
}
