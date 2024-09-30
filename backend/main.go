package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/lockform/pkg/config"
	"github.com/lockform/pkg/database"
	"github.com/lockform/pkg/middleware"
	"github.com/lockform/pkg/routes"
)

func main() {
	database.ConnectDb()
	config.SupertokensInit()
	app := fiber.New(fiber.Config{
		AppName: "Lockform",
	})
	middleware.FiberMiddleware(app)
	routes.InitRoutes(app)

	if err := app.Listen("0.0.0.0:3000"); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
