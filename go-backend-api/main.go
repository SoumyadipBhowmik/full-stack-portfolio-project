package main

import (
	"context"
	"os"

	"github.com/SoumyadipBhowmik/go-backend/driver"
	"github.com/SoumyadipBhowmik/go-backend/route"
	"github.com/SoumyadipBhowmik/go-backend/utils"
	"github.com/gofiber/fiber/v2"
)

var ctx = context.Background()

func init() {
	driver.LoadEnv()
}

func main() {
	app := fiber.New()
	route.InitializeRoutes(app)
	db := driver.ConnectToPostgresDB(ctx)
	defer db.Close()
	port := utils.Checkport(os.Getenv("PORT"))
	err := app.Listen(":" + port)
	utils.CommonErrorCheck(err)

}
