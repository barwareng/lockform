package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lockform/app/models"
	"github.com/lockform/pkg/database"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"gorm.io/gorm"
)

/*
First check if contact of said value and type exists
If does not exist, first create the contact's record. Then associate it with the adding team/individual
If it does exist, associate it with the adding team/individual
*/
func SaveContact(c *fiber.Ctx) error {
	type SaveContactRequest struct {
		Contact     models.Contact     `json:"contact"`
		TeamContact models.TeamContact `json:"teamContact"`
	}
	saveContactRequest := SaveContactRequest{}
	if err := c.BodyParser(&saveContactRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	if err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where(models.Contact{Value: saveContactRequest.Contact.Value, Type: saveContactRequest.Contact.Value}).FirstOrCreate(&saveContactRequest.Contact).Error; err != nil {
			return err
		}
		sessionContainer := session.GetSessionFromRequestContext(c.Context())
		saveContactRequest.TeamContact.ContactID = saveContactRequest.Contact.ID
		saveContactRequest.TeamContact.AddedByID = sessionContainer.GetUserID()
		saveContactRequest.TeamContact.TeamID = c.Locals("teamId").(string)
		if err := tx.Where(models.TeamContact{ContactID: saveContactRequest.Contact.ID, TeamID: saveContactRequest.TeamContact.TeamID}).FirstOrCreate(&saveContactRequest.TeamContact).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
	})
}

func GetContacts(c *fiber.Ctx) error {
	var contacts []models.Contact
	// teamId := c.Locals("teamId").(string)
	// if err := database.DB.Find(&contacts, &models.Contact{TeamID: teamId}).Error; err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   err.Error(),
	// 	})
	// }
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
