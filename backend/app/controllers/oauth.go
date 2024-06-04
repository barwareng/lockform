package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lockform/app/models"
)

func AuthorizationRequest(c *fiber.Ctx) error {
	oauthAuthorizationRequest := new(models.OauthAuthorizationRequest)
	if err := c.QueryParser(oauthAuthorizationRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// sessionTokens := session.RefreshSession()

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"data":  oauthAuthorizationRequest.RedirectUri,
	})
}
