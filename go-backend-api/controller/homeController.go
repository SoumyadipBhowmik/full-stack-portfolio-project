package controller

import "github.com/gofiber/fiber/v2"

type Home struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func ShowHomePage(c *fiber.Ctx) error {
	return c.JSON(Home{Name: "Bruh", Type: "Wooden"})
}
