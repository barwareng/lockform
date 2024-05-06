package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/veriform/app/controllers"
	"github.com/veriform/pkg/middleware"
)

func teamRoutes(app *fiber.App) {
	route := app.Group("/teams", adaptor.HTTPMiddleware(middleware.VerifySession))
	route.Post("/create", controllers.AddTeam)
	route.Get("/", controllers.GetTeams)
}
