package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lockform/app/models"
	"github.com/lockform/pkg/database"
)

type VerificationRequest struct {
	Values []string
}
type VerificationResponse struct {
	TeamMembers []string `json:"teamMembers"`
	Contacts    []string `json:"contacts"`
	Verified    []string `json:"verified"`
	Unknown     []string `json:"unknown"`
	Flagged     []string `json:"flagged"`
}

func VerifyEmailsFromAddon(c *fiber.Ctx) error {
	verificationRequest := &VerificationRequest{}
	var verificationResponse VerificationResponse
	if err := c.BodyParser(verificationRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	var err error
	// teamId := c.Locals("teamId").(string)
	teamId := "cqplou17l2ilf32cufg0"
	verificationRequest.Values, verificationResponse.TeamMembers, err = populateTeamMember(verificationRequest.Values, teamId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	verificationRequest.Values, verificationResponse.Contacts, err = populateContacts(verificationRequest.Values, teamId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	verificationRequest.Values, verificationResponse.Verified, err = populateVerifiedContacts(verificationRequest.Values)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	verificationResponse.Unknown = verificationRequest.Values
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"data":  verificationResponse,
	})
}

// Returns the remaining emails, emails found in team members, and err
func populateTeamMember(emails []string, teamId string) ([]string, []string, error) {
	var members []models.User
	var teamMemberEmails []string
	if err := database.DB.
		Joins("JOIN user_teams ON user_teams.user_id = users.id").
		Where("user_teams.team_id = ? AND users.email IN ?", teamId, emails).
		Find(&members).
		Error; err != nil {
		return nil, nil, err
	}
	for _, member := range members {
		teamMemberEmails = append(teamMemberEmails, member.Email)
	}
	emails = removeElements(emails, teamMemberEmails)
	return emails, teamMemberEmails, nil
}

// Returns the remaining emails, emails found in contacts, and err
func populateContacts(values []string, teamId string) ([]string, []string, error) {
	var contacts []models.Contact
	var contactValues []string
	if err := database.DB.
		Joins("JOIN team_contacts ON team_contacts.contact_id = contacts.id").
		Where("contacts.value IN ? AND team_contacts.team_id =?", values, teamId).
		Find(&contacts).
		Error; err != nil {
		return nil, nil, err
	}
	for _, contact := range contacts {
		contactValues = append(contactValues, contact.Value)
	}
	values = removeElements(values, contactValues)
	return values, contactValues, nil
}

// Returns the remaining emails, emails found to be verified, and err
func populateVerifiedContacts(searchValues []string) ([]string, []string, error) {
	var channels []models.Channel
	var verifiedEmails []string
	if err := database.DB.Where("value IN ?", searchValues).Find(&channels).Error; err != nil {
		return nil, nil, err
	}
	for _, contact := range channels {
		verifiedEmails = append(verifiedEmails, contact.Value)
	}
	searchValues = removeElements(searchValues, verifiedEmails)
	return searchValues, verifiedEmails, nil
}

func removeElements(slice1, slice2 []string) []string {
	// Create a map to store elements of slice2
	removeMap := make(map[string]struct{}, len(slice2))
	for _, v := range slice2 {
		removeMap[v] = struct{}{}
	}

	// Iterate over slice1 and collect elements not in removeMap
	var result []string
	for _, v := range slice1 {
		if _, found := removeMap[v]; !found {
			result = append(result, v)
		}
	}
	return result
}
