package routes

import (
	"github.com/SoumyadipBhowmik/go-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func initializeFeedbackRoutes(app *fiber.App, feedbackController *controllers.UserFeedbackController) {
	api := app.Group("/feedback")
	api.Post("", feedbackController.CreateNewFeedback)
}
