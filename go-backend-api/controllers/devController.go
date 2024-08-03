package controllers

import (
	"github.com/SoumyadipBhowmik/go-backend/repositories"
)

type ContributerController struct {
	contributerRepo *repositories.ContributerRepostory
}

// func AddContributers(c *fiber.Ctx) error {

// 	// contributer := models.DevDTO
// 	// err := c.BodyParser(card)
// 	// if err != nil {
// 		// c.Status(fiber.StatusBadRequest).SendString(err.Error())
// 		// return err
// 	}
// 	// return c.Status(fiber.StatusOK).JSON(card)
// }
