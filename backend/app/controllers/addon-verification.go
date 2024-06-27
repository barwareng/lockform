package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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
	teamId := c.Locals("teamId").(string)
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

// TODO implement team id
func populateTeamMember(searchValues []string, teamId string) (newSearchValues []string, foundTeamMemberEmails []string, err error) {
	var members []models.User
	var teamMemberEmails []string
	log.Info(teamId)
	if err := database.DB.Where("email IN ?", searchValues).Find(&members).Error; err != nil {
		return nil, nil, err
	}
	for _, member := range members {
		teamMemberEmails = append(teamMemberEmails, member.Email)
	}
	searchValues = removeElements(searchValues, teamMemberEmails)
	return searchValues, teamMemberEmails, nil
}
func populateContacts(searchValues []string, teamId string) (newSearchValues []string, foundContactEmails []string, err error) {
	var contacts []models.Contact
	var contactEmails []string
	if err := database.DB.Where("value IN ? AND team_id = ?", searchValues, teamId).Find(&contacts).Error; err != nil {
		return nil, nil, err
	}
	for _, contact := range contacts {
		contactEmails = append(contactEmails, contact.Value)
	}
	searchValues = removeElements(searchValues, contactEmails)
	return searchValues, contactEmails, nil
}
func populateVerifiedContacts(searchValues []string) (newSearchValues []string, foundContactEmails []string, err error) {
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
