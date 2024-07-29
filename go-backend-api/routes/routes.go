package routes

import (
	"github.com/SoumyadipBhowmik/go-backend/controllers"
	"github.com/SoumyadipBhowmik/go-backend/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitializeRoutes(app *fiber.App, db *pgxpool.Pool) {
	repo := repositories.NewUserFeedBackRepository(db)
	controller := controllers.NewUserFeedbackController(repo)

	initializeFeedbackRoutes(app, controller)
}
