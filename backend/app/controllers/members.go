package controllers

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartyemailpassword"
	"github.com/supertokens/supertokens-golang/recipe/userroles"
	"github.com/supertokens/supertokens-golang/recipe/userroles/userrolesclaims"
	"github.com/veriform/app/models"
	"github.com/veriform/pkg/config"
	"github.com/veriform/pkg/database"
)

type AddMemberParams struct {
	UserID string `json:"userId"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}
type Members struct {
	models.User
	Role string `json:"role"`
}

func AddMember(c *fiber.Ctx) error {
	user := &models.User{}
	params := &AddMemberParams{}
	// Check, if received JSON data is valid.
	if err := c.QueryParser(params); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	user.Email = params.Email
	signUpResult, err := thirdpartyemailpassword.EmailPasswordSignUp("public", user.Email, os.Getenv("FAKE_PASSWORD"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// TODO if user already exists just assign team instead of creating new user
	if signUpResult.EmailAlreadyExistsError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   signUpResult.EmailAlreadyExistsError,
		})
	}
	// we successfully created the user. Now we should send them their invite link
	_, err = thirdpartyemailpassword.SendResetPasswordEmail("public", signUpResult.OK.User.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// Save the user in the DB
	teamID := c.Cookies("teamId")
	user.ID = signUpResult.OK.User.ID
	if err := database.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// Associate the use with a team
	if err = database.DB.Model(&user).Association("Teams").Append(&models.Team{ID: teamID}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// TODO Add actual role
	if err = config.AddUserToRole(user.ID, teamID+"_"+params.Role); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// TODO send the user an email letting them know they have been added as ROLE to this TEAM.
	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"member": user,
	})
}
func ChangeMemberRole(c *fiber.Ctx) error {
	params := &AddMemberParams{}
	// Check, if received JSON data is valid.
	if err := c.QueryParser(params); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// Save the user in the DB
	teamId := c.Cookies("teamId")
	currentRolesResp, err := userroles.GetRolesForUser("public", params.UserID, nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	var currentRole string
	for _, role := range currentRolesResp.OK.Roles {
		if strings.HasPrefix(role, teamId+"_") {
			currentRole = role
		}
	}
	sessionContainer := session.GetSessionFromRequestContext(c.Context())
	deleteRoleResp, err := userroles.RemoveUserRole(sessionContainer.GetTenantId(), sessionContainer.GetUserID(), currentRole, nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	if deleteRoleResp.UnknownRoleError != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   deleteRoleResp.UnknownRoleError,
		})
	}

	sessionContainer.FetchAndSetClaim(userrolesclaims.UserRoleClaim)
	sessionContainer.FetchAndSetClaim(userrolesclaims.PermissionClaim)

	if err := config.AddUserToRole(params.UserID, teamId+"_"+params.Role); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	sessionContainer.FetchAndSetClaim(userrolesclaims.UserRoleClaim)
	sessionContainer.FetchAndSetClaim(userrolesclaims.PermissionClaim)
	// TODO send the user an email letting them know they have been added as ROLE to this TEAM.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
	})
}
func GetMembers(c *fiber.Ctx) error {
	teamId := c.Cookies("teamId")
	team := &models.Team{ID: teamId}
	if err := database.DB.Preload("Users").Find(&team).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	var members []Members
	for _, user := range team.Users {
		response, err := userroles.GetRolesForUser("public", user.ID, nil)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
		for _, role := range response.OK.Roles {
			if strings.HasPrefix(role, teamId+"_") {
				member := &Members{User: user, Role: strings.Replace(role, teamId+"_", "", 1)}
				members = append(members, *member)
			}
		}
	}
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"data":  members,
	})
}
