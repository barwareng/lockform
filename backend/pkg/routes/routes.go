package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/veriform/pkg/middleware"
)

func InitRoutes(app *fiber.App) {

	protected := app.Group("/api",
		adaptor.HTTPMiddleware(middleware.VerifySession),
	)
	teamRoutes(protected)
	userRoutes(protected)
	memberRoutes(protected)
	channelRoutes(protected)

	public := app.Group("/public")
	publicChannelRoutes(public)
	app.Get("/badge", func(c *fiber.Ctx) error {
		c.Set("Cross-Origin-Resource-Policy", "same-site")
		return c.Next()
	}, func(c *fiber.Ctx) error {
		// Serve the file using SendFile function
		err := filesystem.SendFile(c, http.Dir(""), "./assets/images/png/verification.png")
		if err != nil {
			// Handle the error, e.g., return a 404 Not Found response
			return c.Status(fiber.StatusNotFound).SendString("File not found")
		}

		return nil
	})

}
