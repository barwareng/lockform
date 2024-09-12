package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lockform/app/models"
	"github.com/lockform/pkg/database"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"gorm.io/gorm"
)

/*
Assign a contact to a team. First, create it if it doesn't exist.
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
		if err := tx.Where(models.Contact{Value: saveContactRequest.Contact.Value, Type: saveContactRequest.Contact.Type}).FirstOrCreate(&saveContactRequest.Contact).Error; err != nil {
			return err
		}
		sessionContainer := session.GetSessionFromRequestContext(c.Context())
		saveContactRequest.TeamContact.AddedByID = sessionContainer.GetUserID()
		saveContactRequest.TeamContact.ContactID = saveContactRequest.Contact.ID
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
		"msg":   "contact saved.",
	})
}

// Find all contacts belonging to a team. Both trusted and untrusted.
func GetTeamContacts(c *fiber.Ctx) error {
	var contacts []models.ContactList
	teamId := c.Locals("teamId").(string)
	if err := database.DB.Table("contacts").
		Select("contacts.id, contacts.value, contacts.label, contacts.domain, contacts.url, contacts.type, contacts.category, team_contacts.is_trusted, team_contacts.reason_for_untrusting, users.id || ',' || users.first_name || ' ' || users.last_name || ',' || users.email AS added_by").
		Joins("JOIN team_contacts ON team_contacts.contact_id = contacts.id").
		Joins("JOIN users ON users.id = team_contacts.added_by_id").
		Where("team_contacts.team_id = ?", teamId).
		Scopes(Paginate(c)).
		Scan(&contacts).Error; err != nil {
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

// Update contact relationship for team. This can be done only by the person that created the contact
func ModifyTrustworthiness(c *fiber.Ctx) error {
	teamContact := models.TeamContact{}
	if err := c.BodyParser(&teamContact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	if err := database.DB.
		Model(&teamContact).Select("IsTrusted", "ReasonForUntrusting").
		Where("contact_id = ?", teamContact.ContactID).
		Updates(teamContact).
		Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "contact updated.",
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
