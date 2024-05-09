package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/veriform/app/models"
	"github.com/veriform/pkg/database"
)

func UpdateProfile(c *fiber.Ctx) error {
	user := &models.User{}
	if err := c.BodyParser(user); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	sessionContainer := session.GetSessionFromRequestContext(c.Context())
	user.ID = sessionContainer.GetUserID()
	user.IsOnboarded = true
	if err := database.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// TODO fix parload not updating on the frontend
	if err := sessionContainer.MergeIntoAccessTokenPayload(map[string]interface{}{"isOnboarded": user.IsOnboarded}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"data":  user,
	})
}

func GetProfile(c *fiber.Ctx) error {
	var user models.User

	sessionContainer := session.GetSessionFromRequestContext(c.Context())
	userID := sessionContainer.GetUserID()
	if err := database.DB.Find(&user, models.User{ID: userID}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"data":  user,
	})
}
