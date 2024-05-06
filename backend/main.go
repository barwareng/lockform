package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/veriform/pkg/config"
	"github.com/veriform/pkg/database"
	"github.com/veriform/pkg/middleware"
	"github.com/veriform/pkg/routes"
)

func main() {
	database.ConnectDb()
	config.SupertokensInit()
	app := fiber.New()
	middleware.FiberMiddleware(app)
	routes.InitRoutes(app)
	if err := app.Listen(":3000"); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
