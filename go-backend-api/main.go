package main

import (
	"context"
	"fmt"
	"os"

	"github.com/SoumyadipBhowmik/go-backend/drivers"
	"github.com/SoumyadipBhowmik/go-backend/routes"
	"github.com/SoumyadipBhowmik/go-backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var ctx = context.Background()

func init() {
	utils.LoadEnv()
}

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("ALLOWED_ORIGINS"),
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))
	db := drivers.ConnectToPostgresDB(ctx)
	routes.InitializeRoutes(app, db)
	defer db.Close()
	port := utils.Checkport(os.Getenv("PORT"))
	err := app.Listen(":" + port)

	if err != nil {
		fmt.Println(err)
	}
}
