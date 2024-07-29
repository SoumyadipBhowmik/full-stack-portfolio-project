package main

import (
	"context"
	"fmt"
	"os"

	"github.com/SoumyadipBhowmik/go-backend/drivers"
	"github.com/SoumyadipBhowmik/go-backend/routes"
	"github.com/SoumyadipBhowmik/go-backend/utils"
	"github.com/gofiber/fiber/v2"
)

var ctx = context.Background()

func init() {
	utils.LoadEnv()
}

func main() {
	app := fiber.New()
	db := drivers.ConnectToPostgresDB(ctx)
	routes.InitializeRoutes(app, db)
	defer db.Close()
	port := utils.Checkport(os.Getenv("PORT"))
	err := app.Listen(":" + port)
	if err != nil {
		fmt.Println(err)
	}
}
