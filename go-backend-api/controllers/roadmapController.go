package controllers

import (
	"github.com/SoumyadipBhowmik/go-backend/models/dto"
	"github.com/SoumyadipBhowmik/go-backend/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type RoadMapController struct {
	service *services.RoadMapService
}

func NewRoadMapController(service *services.RoadMapService) *RoadMapController {
	return &RoadMapController{service: service}
}

func (c *RoadMapController) RoadMapCreation(app *fiber.Ctx) error {
	var roadmap dto.RoadMapDTO
	err := app.BodyParser(&roadmap)
	if err != nil {
		log.Fatal(err.Error())
	}
	return app.JSON(c.service.RoadMapCreation(roadmap.Name, roadmap.Start, roadmap.ExpectedTime, roadmap.Active, roadmap.LeadBy))
}
