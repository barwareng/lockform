package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/lockform/app/controllers"
	"github.com/lockform/pkg/middleware"
)

func adminChannelRoutes(router fiber.Router) {
	route := router.Group("/channels", middleware.ValidateTeam, middleware.ValidateRoles([]string{"owner", "admin"}))
	route.Post("/", middleware.ValidateTeam, controllers.SaveChannel)
	route.Delete("/", controllers.RemoveChannel)
}
func protectedChannelRoutes(router fiber.Router) {
	route := router.Group("/channels", adaptor.HTTPMiddleware(middleware.VerifySession))
	route.Get("/", middleware.ValidateTeam, controllers.GetChannels)
}

func publicChannelRoutes(router fiber.Router) {
	route := router.Group("/verify")
	route.Get("/", controllers.SearchChannel)
}
func channelRoutes(router fiber.Router) {
	// Order of the routes matters. Begin with most accessible
	protectedChannelRoutes(router)
	adminChannelRoutes(router)
}
