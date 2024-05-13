package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/veriform/app/controllers"
	"github.com/veriform/pkg/middleware"
)

func adminChannelRoutes(router fiber.Router) {
	route := router.Group("/channels", middleware.ValidateRoles([]string{"owner", "admin"}))
	route.Post("/", controllers.AddChannel)
	route.Put("/", controllers.AddChannel)
}
func protectedChannelRoutes(router fiber.Router) {
	route := router.Group("/channels", adaptor.HTTPMiddleware(middleware.VerifySession))
	route.Get("/", controllers.GetChannels)
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
