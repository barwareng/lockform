package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lockform/app/controllers"
	"github.com/lockform/pkg/middleware"
)

func adminTeamRoutes(router fiber.Router) {
	route := router.Group("/teams", middleware.ValidateTeam, middleware.ValidateRoles([]string{"owner", "admin"}))
	route.Delete("/", func(c *fiber.Ctx) error { return nil })
}
func protectedTeamRoutes(router fiber.Router) {
	route := router.Group("/teams")
	route.Post("/", controllers.AddTeam)
	route.Get("/", controllers.GetTeams)
}

func teamRoutes(router fiber.Router) {
	// Order of the routes matters. Begin with most accessible
	protectedTeamRoutes(router)
	adminTeamRoutes(router)
}
