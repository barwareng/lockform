package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lockform/app/models"
	"github.com/lockform/pkg/database"
	"github.com/supertokens/supertokens-golang/recipe/session"
)

func SaveTrustedContact(c *fiber.Ctx) error {
	trustedContact := &models.TrustedContact{}
	if err := c.BodyParser(trustedContact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	sessionContainer := session.GetSessionFromRequestContext(c.Context())
	trustedContact.AddedByID = sessionContainer.GetUserID()
	trustedContact.TeamID = c.Locals("teamId").(string)
	if err := database.DB.Save(&trustedContact).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"data":  trustedContact,
	})
}

func GetTrustedContacts(c *fiber.Ctx) error {
	var trustedContacts []models.TrustedContact
	teamId := c.Locals("teamId").(string)
	if err := database.DB.Find(&trustedContacts, &models.TrustedContact{TeamID: teamId}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"data":  trustedContacts,
	})
}

func RemoveTrustedContact(c *fiber.Ctx) error {
	trustedContact := &models.TrustedContact{}
	if err := c.QueryParser(trustedContact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	if err := database.DB.Delete(&trustedContact).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"data":  nil,
	})
}
