package middleware

import "github.com/gofiber/fiber/v2"

func ValidateTeam(c *fiber.Ctx) error {
	if c.Cookies("teamId") != "" {
		c.Locals("teamId", c.Cookies("teamId"))
		return c.Next()
	} else if c.Get("teamId") != "" {
		c.Locals("teamId", c.Get("teamId"))
		return c.Next()
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "team is not selected",
		})
	}
}
