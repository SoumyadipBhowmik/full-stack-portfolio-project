package services

import "github.com/SoumyadipBhowmik/go-backend/repositories"

type ContributerServices struct {
	contributerRepo *repositories.ContributerRepostory
}

func NewContributerServices(contributerRepo *repositories.ContributerRepostory) *ContributerServices {
	return &ContributerServices{contributerRepo: contributerRepo}
}
