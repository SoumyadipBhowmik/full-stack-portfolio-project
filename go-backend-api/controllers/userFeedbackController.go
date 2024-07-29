package controllers

import (
	"fmt"

	"github.com/SoumyadipBhowmik/go-backend/repositories"
	"github.com/gofiber/fiber/v2"
)

type UserFeedbackController struct {
	userRepo *repositories.UserFeedbackRepository
}

func NewUserFeedbackController(repo *repositories.UserFeedbackRepository) *UserFeedbackController {
	return &UserFeedbackController{userRepo: repo}
}

func (fc *UserFeedbackController) CreateNewFeedback(app *fiber.Ctx) error {

	err := fc.userRepo.CreateNewFeedback("Soumyadip", "bhowmik@gmail.com", "Wow, SO freaking Awesome Bruh")
	if err != nil {
		fmt.Println("Beuh")
	}
	return nil
}
