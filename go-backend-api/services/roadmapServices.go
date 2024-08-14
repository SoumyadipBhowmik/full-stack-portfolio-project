package services

import (
	"time"

	"github.com/SoumyadipBhowmik/go-backend/models/dto"
	"github.com/SoumyadipBhowmik/go-backend/repositories"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jinzhu/copier"
)

type RoadMapService struct {
	repo *repositories.RoadMapRepository
}

func NewRoadMapService(repo *repositories.RoadMapRepository) *RoadMapService {
	return &RoadMapService{repo: repo}
}

func (service *RoadMapService) RoadMapCreation(name string, startDate, expectedEndDate time.Time, active bool, leadBy string) *dto.RoadMapDTO {

	var roadMapDTO dto.RoadMapDTO
	roadmap := service.repo.RoadMapCreation(name, startDate, expectedEndDate, active, leadBy)
	err := copier.Copy(&roadMapDTO, roadmap)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &roadMapDTO
}
