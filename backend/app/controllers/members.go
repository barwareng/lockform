package controllers

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartyemailpassword"
	"github.com/supertokens/supertokens-golang/recipe/userroles"
	"github.com/veriform/app/models"
	"github.com/veriform/pkg/config"
	"github.com/veriform/pkg/database"
)

type AddMemberParams struct {
	Email string `json:"email"`
	Role  string `json:"role"`
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
	// Assoaciate the use with a team
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
func GetMembers(c *fiber.Ctx) error {
	roles, err := userroles.GetAllRoles(nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	var userIds []string
	for _, role := range roles.OK.Roles {
		response, err := userroles.GetUsersThatHaveRole("public", role, nil)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		if response.UnknownRoleError != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   "Tried to get users from a non-existent team role",
			})
		}
		userIds = append(userIds, response.OK.Users...)
	}
	users := []models.User{}
	result := database.DB.Find(&users, userIds)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   result.Error.Error(),
		})
	}
	return c.JSON(users)
}
