package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/lockform/app/controllers"
	"github.com/lockform/pkg/middleware"
)

func adminContactRoutes(router fiber.Router) {
	route := router.Group("/contacts", middleware.ValidateTeam, middleware.ValidateRoles([]string{"owner", "admin"}))
	route.Delete("/", controllers.RemoveContact)
}
func protectedContactRoutes(router fiber.Router) {
	route := router.Group("/contacts", middleware.ValidateTeam, adaptor.HTTPMiddleware(middleware.VerifySession))
	route.Post("/", controllers.SaveContact)
	route.Get("/", controllers.GetTeamContacts)
	route.Put("/", controllers.ModifyTrustworthiness)
}
func contactRoutes(router fiber.Router) {
	// Order of the routes matters. Begin with most accessible
	protectedContactRoutes(router)
	adminContactRoutes(router)
}
