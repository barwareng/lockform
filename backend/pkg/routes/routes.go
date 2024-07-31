package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/lockform/app/controllers"
	"github.com/lockform/pkg/middleware"
)

func InitRoutes(app *fiber.App) {

	protected := app.Group("/api",
		adaptor.HTTPMiddleware(middleware.VerifySession),
	)
	// OauthRoutes
	protected.Post("/oauth/authorize", controllers.AuthorizationRequest)
	// Protected Resource Routes
	teamRoutes(protected)
	userRoutes(protected)
	memberRoutes(protected)
	channelRoutes(protected)
	contactRoutes(protected)
	protected.Post("/verification", middleware.ValidateTeam, controllers.VerifyEmailsFromAddon)
	// Public Channels
	public := app.Group("/public")
	public.Post("/verification", controllers.VerifyEmailsFromAddon)
	publicChannelRoutes(public)

}
