package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/veriform/app/controllers"
)

func protectedUserRoutes(router fiber.Router) {
	route := router.Group("/users")
	route.Put("/", controllers.UpdateProfile)
	route.Get("/", controllers.GetProfile)
}

func userRoutes(router fiber.Router) {
	// Order of the routes matters. Begin with most accessible
	protectedUserRoutes(router)
}
