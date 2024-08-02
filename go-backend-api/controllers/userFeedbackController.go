package controllers

import (
	"fmt"

	models "github.com/SoumyadipBhowmik/go-backend/models/dto"
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

	var user models.UserDTO

	err := app.BodyParser(&user)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid body request")
	}
	feedback, err := fc.userService.CreateNewFeedback(user.Name, user.Email, user.Feedback)
	if err != nil {
		fmt.Println("error occured:", err)
	}
	fmt.Println(app.JSON(feedback))
	return app.JSON(feedback)
}
