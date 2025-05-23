package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/lockform/app/controllers"
	"github.com/lockform/pkg/middleware"
)

func adminMemberRoutes(router fiber.Router) {
	route := router.Group("/members", middleware.ValidateTeam, middleware.ValidateRoles([]string{"owner", "admin"}))
	route.Post("/", controllers.AddMember)
	route.Put("/", controllers.ChangeMemberRole)
	route.Delete("/", controllers.RemoveMember)
}
func protectedMemberRoutes(router fiber.Router) {
	route := router.Group("/members", adaptor.HTTPMiddleware(middleware.VerifySession))
	route.Get("/", middleware.ValidateTeam, controllers.GetMembers)
}
func memberRoutes(router fiber.Router) {
	// Order of the routes matters. Begin with most accessible
	protectedMemberRoutes(router)
	adminMemberRoutes(router)
}
