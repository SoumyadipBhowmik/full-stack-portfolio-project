package controllers

import "github.com/gofiber/fiber/v2"

type Home struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func ShowHomeCards(c *fiber.Ctx) error {
	return c.JSON(Home{Name: "Bruh", Type: "Wooden"})
}

func CreateCards(c *fiber.Ctx) error {

	card := new(Home)
	err := c.BodyParser(card)
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}
	return c.Status(fiber.StatusOK).JSON(card)
}
