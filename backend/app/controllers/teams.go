package controllers

import (
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/userroles"
	"github.com/supertokens/supertokens-golang/recipe/userroles/userrolesclaims"
	"github.com/veriform/app/models"
	"github.com/veriform/pkg/database"
	"gorm.io/gorm"
)

func AddTeam(c *fiber.Ctx) error {
	sessionContainer := session.GetSessionFromRequestContext(c.Context())
	userID := sessionContainer.GetUserID()
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
	if err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&team).Error; err != nil {
			return err
		}
		roles := []string{"owner", "admin", "developer", "billing", "viewer"}
		for _, role := range roles {
			_, err := userroles.CreateNewRoleOrAddPermissions(team.ID+"_"+role, []string{
				"read",
			}, nil)
			if err != nil {
				return err
			}
		}
		if err := tx.Model(&team).Association("Users").Append(&models.User{ID: userID}); err != nil {
			return err
		}
		response, err := userroles.AddRoleToUser("public", userID, team.ID+"_owner", nil)
		if err != nil {
			return err
		}

		if response.UnknownRoleError != nil {
			return errors.New("unknown error trying to add role to user")
		}
		if err := sessionContainer.FetchAndSetClaim(userrolesclaims.UserRoleClaim); err != nil {
			return err
		}
		var userTeams []models.AccessTokenTeamPayload
		accessTokenPayload := sessionContainer.GetAccessTokenPayload()
		if accessTokenPayload["teams"] != nil {
			teams, ok := accessTokenPayload["teams"].([]interface{})
			if !ok {
				return errors.New("could not assert claim")
			}
			for _, team := range teams {
				teamJson, err := json.Marshal(team)
				if err != nil {
					return err
				}
				var payload models.AccessTokenTeamPayload
				err = json.Unmarshal(teamJson, &payload)
				if err != nil {
					return err
				}
				userTeams = append(userTeams, payload)
			}
		}
		userTeams = append(userTeams, models.AccessTokenTeamPayload{ID: team.ID, Name: team.Name})
		accessTokenPayload["teams"] = userTeams
		if err := sessionContainer.MergeIntoAccessTokenPayload(accessTokenPayload); err != nil {
			return err
		}
		// return errors.New("just to cap")
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
		"data":  team,
	})
}

func GetTeams(c *fiber.Ctx) error {
	teams := []models.Team{}
	if err := database.DB.Preload("Users").Find(&teams).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"data":  teams,
	})
}
