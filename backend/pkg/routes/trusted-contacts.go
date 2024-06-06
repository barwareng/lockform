package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/lockform/app/controllers"
	"github.com/lockform/pkg/middleware"
)

func adminTrustedContactRoutes(router fiber.Router) {
	route := router.Group("/trusted-contacts", middleware.ValidateTeam, middleware.ValidateRoles([]string{"owner", "admin"}))
	route.Delete("/", controllers.RemoveTrustedContact)
}
func protectedTrustedContactRoutes(router fiber.Router) {
	route := router.Group("/trusted-contacts", middleware.ValidateTeam, adaptor.HTTPMiddleware(middleware.VerifySession))
	route.Post("/", controllers.SaveTrustedContact)
	route.Get("/", controllers.GetChannels)
}
func trustedContactRoutes(router fiber.Router) {
	// Order of the routes matters. Begin with most accessible
	protectedTrustedContactRoutes(router)
	adminTrustedContactRoutes(router)
}
