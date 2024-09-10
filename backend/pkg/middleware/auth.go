package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/userroles/userrolesclaims"
)

func VerifySession(handler http.Handler) http.Handler {
	return session.VerifySession(nil, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}))
}

func ValidateRoles(roles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		verified, err := verifyRoles(roles, c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
		if !verified {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": true,
				"msg":   "Insufficient permissions",
			})
		}
		return c.Next()

	}
}

func verifyRoles(userRoles []string, c *fiber.Ctx) (bool, error) {
	teamId := c.Locals("teamId").(string)

	var formattedUserRoles []interface{} //Roles are formatted as teamId_role
	for _, role := range userRoles {
		formattedUserRoles = append(formattedUserRoles, teamId+"_"+role)
	}
	sessionContainer := session.GetSessionFromRequestContext(c.Context())
	roles := sessionContainer.GetClaimValue(userrolesclaims.UserRoleClaim)

	if roles == nil || !containsRoles(roles.([]interface{}), formattedUserRoles) {
		return false, nil
	}
	return true, nil
}

func containsRoles(userRoles []interface{}, requiredRoles []interface{}) bool {
	for _, ur := range userRoles {
		for _, rr := range requiredRoles {
			if ur == rr {
				return true
			}
		}
	}
	return false
}
