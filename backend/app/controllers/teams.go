package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/rs/xid"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/userroles"
	"github.com/veriform/app/models"
	"github.com/veriform/pkg/database"
)

func AddTeam(c *fiber.Ctx) error {
	team := &models.Team{}
	// Check, if received JSON data is valid.
	if err := c.BodyParser(team); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	team.ID = xid.New().String()
	if err := database.DB.Create(&team).Error; err != nil {
		log.Info("Creating team failed: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	roles := []string{"owner", "admin", "developer", "billing", "viewer"}
	for _, role := range roles {
		_, err := userroles.CreateNewRoleOrAddPermissions(team.ID+"_"+role, []string{
			"read",
		}, nil)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
	}
	userID := c.GetReqHeaders()["X-User"][0]
	err := database.DB.Model(&team).Association("Users").Append(&models.User{ID: userID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	response, err := userroles.AddRoleToUser("public", userID, team.ID+"_owner", nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if response.UnknownRoleError != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   response.UnknownRoleError,
		})

	}

	sessionContainer := session.GetSessionFromRequestContext(c.Context())
	accessTokenPayload := sessionContainer.GetAccessTokenPayload()
	teams := accessTokenPayload["teams"].([]interface{})
	newTeam := map[string]interface{}{"id": team.ID, "name": team.Name}
	teams = append(teams, newTeam)
	if err := sessionContainer.MergeIntoAccessTokenPayload(map[string]interface{}{"teams": teams}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"data":  team,
	})
}
func GetTeams(c *fiber.Ctx) error {
	log.Info(c.GetReqHeaders()["X-User"])
	teams := []models.Team{}
	database.DB.Preload("Users").Find(&teams)
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"data":  teams,
	})
}
