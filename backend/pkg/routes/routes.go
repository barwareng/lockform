package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/veriform/pkg/middleware"
)

func InitRoutes(app *fiber.App) {
	protected := app.Group("/api", adaptor.HTTPMiddleware(middleware.VerifySession))
	teamRoutes(protected)
	userRoutes(protected)
	memberRoutes(protected)
	channelRoutes(protected)

	public := app.Group("/public")
	publicChannelRoutes(public)
}
