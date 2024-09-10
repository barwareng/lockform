package controllers

import (
	"errors"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/lockform/app/models"
	"github.com/lockform/app/services"
	"github.com/lockform/pkg/database"
	"github.com/rs/xid"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/userroles"
	"github.com/supertokens/supertokens-golang/recipe/userroles/userrolesclaims"
	"gorm.io/gorm"
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
	teamID := c.Locals("teamId").(string)
	user.Email = params.Email
	if err := database.DB.Transaction(func(tx *gorm.DB) error {
		// Check if user already exists
		signedUpUser, err := emailpassword.GetUserByEmail("public", params.Email)
		if err != nil {
			return err
		}

		var invitationLink string
		if signedUpUser != nil {
			// If user already exists, we don't need to create a new user
			user.ID = signedUpUser.ID
			invitationLink = os.Getenv("SUPERTOKENS_FRONTEND_DOMAIN")
		} else {
			// If user does not exist, create a new user
			signUpResult, err := emailpassword.SignUp("public", params.Email, xid.New().String())
			if err != nil {
				return err
			}
			// we successfully created the user. Create a password reset link to send as an invitation link
			link, err := emailpassword.CreateResetPasswordLink("public", signUpResult.OK.User.ID)
			if err != nil {
				return err
			}
			if link.UnknownUserIdError != nil {
				return errors.New("uknown error creating password reset link for user")
			}
			if link.OK.Link != "" {
				invitationLink = strings.Replace(
					link.OK.Link,
					"auth/reset-password",
					"reset-password/new", 1,
				)
			}
			user.ID = signUpResult.OK.User.ID
			if err := tx.Save(&user).Error; err != nil {
				return err
			}
		}
		// Get the team being added to
		team := models.Team{}
		if err := tx.Select("name").Where("id = ?", teamID).Limit(1).Find(&team).Error; err != nil {
			return err
		}
		// Get who is adding the user
		sessionContainer := session.GetSessionFromRequestContext(c.Context())
		addedById := sessionContainer.GetUserID()
		addedBy := models.User{}
		if err := tx.Select("first_name", "email").Where("id = ?", addedById).Limit(1).Find(&addedBy).Error; err != nil {
			return err
		}
		// Associate the use with a team
		if err = tx.Model(&user).Association("Teams").Append(&models.Team{ID: teamID}); err != nil {
			return err
		}
		// Add actual role
		if err = addUserToRole(user.ID, teamID+"_"+params.Role); err != nil {
			return err
		}

		data := models.InviteTeamUser{
			Email:            user.Email,
			InviterFirstName: addedBy.FirstName,
			InviterEmail:     addedBy.Email,
			TeamName:         team.Name,
			InvitationLink:   invitationLink,
		}
		if err := services.SendTeamInvitationEmail(data); err != nil {
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

	// if err := services.SendTeamInvitationEmail(params.Email); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   err.Error(),
	// 	})
	// }
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
