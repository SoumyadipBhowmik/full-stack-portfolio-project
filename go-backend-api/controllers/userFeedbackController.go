package controllers

import (
	"fmt"

	"github.com/SoumyadipBhowmik/go-backend/services"
	"github.com/gofiber/fiber/v2"
)

type UserFeedbackController struct {
	userService *services.UserFeedbackService
}

func NewUserFeedbackController(service *services.UserFeedbackService) *UserFeedbackController {
	return &UserFeedbackController{userService: service}
}

func (fc *UserFeedbackController) CreateNewFeedback(app *fiber.Ctx) error {
	name := "Soumyadip"
	email := "bhowmik.soumyadip9@gmail.com"
	myfeedback := "truely a wonderful experience bruh"
	feedback, err := fc.userService.CreateNewFeedback(name, email, myfeedback)
	if err != nil {
		fmt.Println("error occured:", err)
	}
	fmt.Println(app.JSON(feedback))
	return app.JSON(feedback)
}
