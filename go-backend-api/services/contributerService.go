package services

import (
	"github.com/SoumyadipBhowmik/go-backend/models/dto"
	"github.com/SoumyadipBhowmik/go-backend/repositories"
	"github.com/jinzhu/copier"
)

type ContributerServices struct {
	contributerRepo *repositories.ContributerRepostory
}

func NewContributerServices(contributerRepo *repositories.ContributerRepostory) *ContributerServices {
	return &ContributerServices{contributerRepo: contributerRepo}
}

func (service *ContributerServices) AddContributor(name, role, description, github string) *dto.DevDTO {
	contributor := service.contributerRepo.AddContributor(name, role, description, github)

	var devDTO dto.DevDTO
	copier.Copy(&devDTO, contributor)
	return &devDTO
}
