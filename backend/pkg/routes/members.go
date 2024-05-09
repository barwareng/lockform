package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/veriform/app/controllers"
	"github.com/veriform/pkg/middleware"
)

func adminMemberRoutes(router fiber.Router) {
	route := router.Group("/members", middleware.ValidateRoles([]string{"admin"}))
	route.Post("/", controllers.AddMember)
	route.Put("/", controllers.ChangeMemberRole)
}
func protectedMemberRoutes(router fiber.Router) {
	route := router.Group("/members", adaptor.HTTPMiddleware(middleware.VerifySession))
	route.Get("/", controllers.GetMembers)
}
func memberRoutes(router fiber.Router) {
	// Order of the routes matters. Begin with most accessible
	protectedMemberRoutes(router)
	adminMemberRoutes(router)
}
