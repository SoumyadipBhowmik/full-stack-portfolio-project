package controllers

import (
	"log"

	"github.com/SoumyadipBhowmik/go-backend/models/dto"
	"github.com/SoumyadipBhowmik/go-backend/services"
	"github.com/gofiber/fiber/v2"
)

type ContributerController struct {
	contributerService *services.ContributerServices
}

func NewContributerController(contibuterService *services.ContributerServices) *ContributerController {
	return &ContributerController{contributerService: contibuterService}
}

func (c *ContributerController) AddContributer(app *fiber.Ctx) error {
	var contributor dto.DevDTO
	err := app.BodyParser(&contributor)
	if err != nil {
		log.Fatal(err.Error())
	}
	contribute := c.contributerService.AddContributor(contributor.Name, contributor.Role, contributor.Description, contributor.Github)
	return app.JSON(contribute)
}
