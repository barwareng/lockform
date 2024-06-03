package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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
	log.Info(oauthAuthorizationRequest)
	// return c.Redirect(oauthAuthorizationRequest.RedirectUri)
	return c.Redirect("http://127.0.0.1:5175/callback", 301)
	// return c.SendStatus(200)

}
