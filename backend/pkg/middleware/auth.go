package middleware

import (
	"net/http"

	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/session/claims"
	"github.com/supertokens/supertokens-golang/recipe/session/sessmodels"
	"github.com/supertokens/supertokens-golang/recipe/userroles/userrolesclaims"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func VerifySession(handler http.Handler) http.Handler {
	return session.VerifySession(nil, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}))
}

// Admin role
func VerifyAdmin(handler http.Handler) http.Handler {
	return session.VerifySession(&sessmodels.VerifySessionOptions{
		OverrideGlobalClaimValidators: func(globalClaimValidators []claims.SessionClaimValidator, sessionContainer sessmodels.SessionContainer, userContext supertokens.UserContext) ([]claims.SessionClaimValidator, error) {
			request := supertokens.GetRequestFromUserContext(userContext)
			teamID, err := request.Cookie("teamId")
			if err != nil {
				return nil, err
			}
			roles := []string{"owner", "admin"}
			for _, role := range roles {
				globalClaimValidators = append(globalClaimValidators, userrolesclaims.UserRoleClaimValidators.Includes(teamID.Value+"_"+role, nil, nil))
			}
			return globalClaimValidators, nil
		},
	}, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}))
}

// Super admin role
func VerifyOwner(handler http.Handler) http.Handler {
	return session.VerifySession(&sessmodels.VerifySessionOptions{
		OverrideGlobalClaimValidators: func(globalClaimValidators []claims.SessionClaimValidator, sessionContainer sessmodels.SessionContainer, userContext supertokens.UserContext) ([]claims.SessionClaimValidator, error) {
			request := supertokens.GetRequestFromUserContext(userContext)
			teamID, err := request.Cookie("teamId")
			if err != nil {
				return nil, err
			}
			globalClaimValidators = append(globalClaimValidators, userrolesclaims.UserRoleClaimValidators.Includes(teamID.Value+"_owner", nil, nil))
			return globalClaimValidators, nil
		},
	}, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}))
}
