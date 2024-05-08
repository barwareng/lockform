package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/veriform/app/controllers"
	"github.com/veriform/pkg/middleware"
)

func adminTeamRoutes(router fiber.Router) {
	route := router.Group("/teams/admin", adaptor.HTTPMiddleware(middleware.VerifyAdmin))
	route.Delete("/", func(c *fiber.Ctx) error { return nil })
}
func protectedTeamRoutes(router fiber.Router) {
	route := router.Group("/teams", adaptor.HTTPMiddleware(middleware.VerifySession))
	route.Post("/", controllers.AddTeam)
	route.Get("/", controllers.GetTeams)
}

func teamRoutes(router fiber.Router) {
	adminTeamRoutes(router)
	protectedTeamRoutes(router)
}
