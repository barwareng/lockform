package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lockform/app/models"
	"github.com/lockform/pkg/database"
	"github.com/supertokens/supertokens-golang/recipe/session"
)

func SaveContact(c *fiber.Ctx) error {
	contact := &models.Contact{}
	if err := c.BodyParser(contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	sessionContainer := session.GetSessionFromRequestContext(c.Context())
	contact.AddedByID = sessionContainer.GetUserID()
	contact.TeamID = c.Locals("teamId").(string)
	if err := database.DB.Save(&contact).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"data":  contact,
	})
}

func GetContacts(c *fiber.Ctx) error {
	var contacts []models.Contact
	teamId := c.Locals("teamId").(string)
	if err := database.DB.Find(&contacts, &models.Contact{TeamID: teamId}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"data":  contacts,
	})
}

func RemoveContact(c *fiber.Ctx) error {
	contact := &models.Contact{}
	if err := c.QueryParser(contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	if err := database.DB.Delete(&contact).Error; err != nil {
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
