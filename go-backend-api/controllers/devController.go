package controllers

import (
	"github.com/SoumyadipBhowmik/go-backend/services"
	"github.com/gofiber/fiber/v2"
)

type ContributerController struct {
	contributerService *services.ContributerServices
}

func NewContributerController(contibuterService *services.ContributerServices) *ContributerController {
	return &ContributerController{contributerService: contibuterService}
}

func (c *ContributerController) AddContributer(app *fiber.Ctx) {

}
