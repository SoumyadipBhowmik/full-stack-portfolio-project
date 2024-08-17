package routes

import (
	"github.com/SoumyadipBhowmik/go-backend/controllers"
	"github.com/SoumyadipBhowmik/go-backend/repositories"
	"github.com/SoumyadipBhowmik/go-backend/services"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitializeRoutes(app *fiber.App, db *pgxpool.Pool) {
	feedbackRepo := repositories.NewUserFeedBackRepository(db)
	feedbackService := services.NewUserFeedbackService(feedbackRepo)
	feedbackController := controllers.NewUserFeedbackController(feedbackService)

	initializeFeedbackRoutes(app, feedbackController)

	devRepo := repositories.NewContributerRepository(db)
	devService := services.NewContributerServices(devRepo)
	devController := controllers.NewContributerController(devService)

	initializeDevRoutes(app, devController)

	roadmapRepo := repositories.NewRoadMapRepository(db)
	roadmapService := services.NewRoadMapService(roadmapRepo)
	roadmapController := controllers.NewRoadMapController(roadmapService)

	initializeRoadMapRoutes(app, roadmapController)

	postRepo := repositories.NewPostRepository(db)
	postService := services.NewPostServices(postRepo)
	postControntroller := controllers.NewPostController(postService)

	initializePostRoutes(app, postControntroller)

}
