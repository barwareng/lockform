package controllers

import (
	"errors"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/lockform/app/models"
	"github.com/lockform/app/services"
	"github.com/lockform/pkg/database"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartyemailpassword"
	"github.com/supertokens/supertokens-golang/recipe/userroles"
	"github.com/supertokens/supertokens-golang/recipe/userroles/userrolesclaims"
)

type AddMemberParams struct {
	UserID string `json:"userId:omit_empty"`
	Email  string `json:"email:omit_empty"`
	Role   string `json:"role"`
}
type Members struct {
	models.User
	Role string `json:"role"`
}

// TODO use transactions
func AddMember(c *fiber.Ctx) error {
	var user models.User
	params := &AddMemberParams{}
	// Check, if received JSON data is valid.
	if err := c.QueryParser(params); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// Check if user already exists
	getUsersResult, err := thirdpartyemailpassword.GetUsersByEmail("public", params.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// Save the user in the DB
	teamID := c.Locals("teamId").(string)
	user.Email = params.Email
	if len(getUsersResult) == 0 {
		signUpResult, err := thirdpartyemailpassword.EmailPasswordSignUp("public", params.Email, os.Getenv("FAKE_PASSWORD"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
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
		user.ID = signUpResult.OK.User.ID
		if err := database.DB.Save(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
	} else {
		user.ID = getUsersResult[0].ID

	}

	// Associate the use with a team
	if err = database.DB.Model(&user).Association("Teams").Append(&models.Team{ID: teamID}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// Add actual role
	if err = addUserToRole(user.ID, teamID+"_"+params.Role); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	if err := services.SendTeamInvitationEmail(params.Email); err != nil {
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
func ChangeMemberRole(c *fiber.Ctx) error {
	params := &AddMemberParams{}
	if err := c.QueryParser(params); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	teamId := c.Locals("teamId").(string)
	if err := removeUserRole(params.UserID, teamId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	if err := addUserToRole(params.UserID, teamId+"_"+params.Role); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	sessionContainer := session.GetSessionFromRequestContext(c.Context())
	sessionContainer.FetchAndSetClaim(userrolesclaims.UserRoleClaim)
	sessionContainer.FetchAndSetClaim(userrolesclaims.PermissionClaim)

	if err := services.SendTeamInvitationEmail(params.Email); err != nil {
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
func GetMembers(c *fiber.Ctx) error {
	teamId := c.Locals("teamId").(string)
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

func RemoveMember(c *fiber.Ctx) error {
	// delete role
	params := &AddMemberParams{}
	if err := c.QueryParser(params); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	teamId := c.Locals("teamId").(string)
	if err := removeUserRole(params.UserID, teamId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// remove association
	team := &models.Team{ID: teamId}
	if err := database.DB.Model(&team).Association("Users").Delete(&models.User{ID: params.UserID}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"data":  "Member removed successfully",
	})
}

func removeUserRole(userId string, teamId string) error {
	currentRolesResp, err := userroles.GetRolesForUser("public", userId, nil)
	if err != nil {
		return err
	}
	var currentRole string
	for _, role := range currentRolesResp.OK.Roles {
		if strings.HasPrefix(role, teamId+"_") {
			currentRole = role
		}
	}

	deleteRoleResp, err := userroles.RemoveUserRole("public", userId, currentRole, nil)
	if err != nil {
		return err
	}

	if deleteRoleResp.UnknownRoleError != nil {
		return errors.New("uknown error removing role from user")
	}
	return nil
}

func addUserToRole(userID string, role string) error {
	response, err := userroles.AddRoleToUser("public", userID, role, nil)
	if err != nil {
		return err
	}

	if response.UnknownRoleError != nil {
		// No such role exists
		return errors.New("uknown error assigning role to user")
	}

	if response.OK.DidUserAlreadyHaveRole {
		return errors.New("user already exists for this role")
	}
	return nil
}
