package config

import (
	"net/http"
	"os"
	"strings"

	"github.com/lockform/app/models"
	"github.com/lockform/app/services"
	"github.com/lockform/pkg/database"
	"github.com/rs/xid"
	"github.com/supertokens/supertokens-golang/ingredients/emaildelivery"
	"github.com/supertokens/supertokens-golang/recipe/emailverification"
	"github.com/supertokens/supertokens-golang/recipe/emailverification/evmodels"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/session/sessmodels"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty/tpmodels"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartyemailpassword"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartyemailpassword/tpepmodels"
	"github.com/supertokens/supertokens-golang/recipe/userroles"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func SupertokensInit() {
	cookieDomain := os.Getenv("SUPERTOKENS_COOKIE_DOMAIN")
	cookieSecure := true
	cookieSameSite := "lax"
	apiBasePath := "/auth"
	websiteBasePath := "/auth"
	err := supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			ConnectionURI: os.Getenv("SUPERTOKENS_CONNECTION_URI"),
		},
		AppInfo: supertokens.AppInfo{
			AppName:   os.Getenv("APP_NAME"),
			APIDomain: os.Getenv("SUPERTOKENS_API_DOMAIN"),
			// WebsiteDomain:   os.Getenv("SUPERTOKENS_FRONTEND_DOMAIN"),
			GetOrigin: func(request *http.Request, userContext supertokens.UserContext) (string, error) {
				if request != nil {
					origin := request.Header.Get("origin")
					if origin == "" {
						// this means the client is in an iframe, it's a mobile app, or
						// there is a privacy setting on the frontend which doesn't send
						// the origin
					} else {
						if origin == "https://script.google.com" {
							// query from the test site
							return "https://script.google.com", nil
						}
					}
				}
				return os.Getenv("SUPERTOKENS_FRONTEND_DOMAIN"), nil
			},
			APIBasePath:     &apiBasePath,
			WebsiteBasePath: &websiteBasePath,
		},
		RecipeList: []supertokens.Recipe{
			emailverification.Init(evmodels.TypeInput{
				Mode: evmodels.ModeOptional,
				Override: &evmodels.OverrideStruct{
					APIs: func(originalImplementation evmodels.APIInterface) evmodels.APIInterface {
						ogVerifyEmailPOST := *originalImplementation.VerifyEmailPOST
						(*originalImplementation.VerifyEmailPOST) = func(token string, sessionContainer sessmodels.SessionContainer, tenantId string, options evmodels.APIOptions, userContext supertokens.UserContext) (evmodels.VerifyEmailPOSTResponse, error) {
							resp, err := ogVerifyEmailPOST(token, sessionContainer, tenantId, options, userContext)
							if err != nil {
								return evmodels.VerifyEmailPOSTResponse{}, err
							}

							if resp.OK != nil {
								id := resp.OK.User.ID
								email := resp.OK.User.Email
								user := models.User{ID: id, Email: email}
								database.DB.Create(&user)
							}

							return resp, nil
						}
						return originalImplementation
					},
				},
				EmailDelivery: &emaildelivery.TypeInput{
					Override: func(originalImplementation emaildelivery.EmailDeliveryInterface) emaildelivery.EmailDeliveryInterface {

						(*originalImplementation.SendEmail) = func(input emaildelivery.EmailType, userContext supertokens.UserContext) error {

							input.EmailVerification.EmailVerifyLink = strings.Replace(
								input.EmailVerification.EmailVerifyLink,
								"auth/",
								"", 1,
							)
							return services.SendEmailVerificationLink(input.EmailVerification.User.Email, input.EmailVerification.EmailVerifyLink)
						}
						return originalImplementation
					},
				},
			}),
			thirdpartyemailpassword.Init(&tpepmodels.TypeInput{
				Override: &tpepmodels.OverrideStruct{
					Functions: func(originalImplementation tpepmodels.RecipeInterface) tpepmodels.RecipeInterface {
						// create a copy of the originalImplementation
						originalEmailPasswordSignUp := *originalImplementation.EmailPasswordSignUp
						originalThirdPartySignInUp := *originalImplementation.ThirdPartySignInUp

						// override the email password sign up function
						(*originalImplementation.EmailPasswordSignUp) = func(email, password string, tenantId string, userContext supertokens.UserContext) (tpepmodels.SignUpResponse, error) {

							resp, err := originalEmailPasswordSignUp(email, password, tenantId, userContext)
							if err != nil {
								return tpepmodels.SignUpResponse{}, err
							}

							if resp.OK != nil {
								externalUserId := xid.New().String()
								_, err := supertokens.CreateUserIdMapping(resp.OK.User.ID, externalUserId, nil, nil)
								if err != nil {
									return tpepmodels.SignUpResponse{}, err
								}
								resp.OK.User.ID = externalUserId
							}

							return resp, err
						}

						// override the thirdparty sign in / up function
						(*originalImplementation.ThirdPartySignInUp) = func(thirdPartyID, thirdPartyUserID, email string, oAuthTokens tpmodels.TypeOAuthTokens, rawUserInfoFromProvider tpmodels.TypeRawUserInfoFromProvider, tenantId string, userContext supertokens.UserContext) (tpepmodels.SignInUpResponse, error) {

							resp, err := originalThirdPartySignInUp(thirdPartyID, thirdPartyUserID, email, oAuthTokens, rawUserInfoFromProvider, tenantId, userContext)
							if err != nil {
								return tpepmodels.SignInUpResponse{}, err
							}

							if resp.OK != nil {
								if resp.OK.CreatedNewUser {
									externalUserId := xid.New().String()
									_, err := supertokens.CreateUserIdMapping(resp.OK.User.ID, externalUserId, nil, nil)
									if err != nil {
										return tpepmodels.SignInUpResponse{}, err
									}
									resp.OK.User.ID = externalUserId
									user := models.User{ID: externalUserId, Email: email}
									database.DB.Create(&user)
								}
							}

							return resp, err
						}

						return originalImplementation
					},
				},
				EmailDelivery: &emaildelivery.TypeInput{
					Override: func(originalImplementation emaildelivery.EmailDeliveryInterface) emaildelivery.EmailDeliveryInterface {

						(*originalImplementation.SendEmail) = func(input emaildelivery.EmailType, userContext supertokens.UserContext) error {
							input.PasswordReset.PasswordResetLink = strings.Replace(
								input.PasswordReset.PasswordResetLink,
								"auth/reset-password",
								"reset-password/new", 1,
							)
							return services.SendPasswordResetLink(input.PasswordReset.User.Email, input.PasswordReset.PasswordResetLink)
						}
						return originalImplementation
					},
				},
				Providers: []tpmodels.ProviderInput{
					{
						Config: tpmodels.ProviderConfig{
							ThirdPartyId: "google",
							Clients: []tpmodels.ProviderClientConfig{
								{
									ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
									ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
								},
							},
						},
					},
				},
			}),
			userroles.Init(nil),
			session.Init(&sessmodels.TypeInput{
				Override: &sessmodels.OverrideStruct{
					Functions: func(originalImplementation sessmodels.RecipeInterface) sessmodels.RecipeInterface {

						// first we copy the original implementation
						originalCreateNewSession := *originalImplementation.CreateNewSession

						(*originalImplementation.CreateNewSession) = func(userID string, accessTokenPayload, sessionDataInDatabase map[string]interface{}, disableAntiCsrf *bool, tenantId string, userContext supertokens.UserContext) (sessmodels.SessionContainer, error) {
							user := &models.User{ID: userID}
							var userTeams []models.AccessTokenTeamPayload
							if err := database.DB.Preload("Teams").Find(&user).Error; err != nil {
								return nil, err
							}
							if len(user.Teams) > 0 {
								for _, team := range user.Teams {
									userTeams = append(userTeams, models.AccessTokenTeamPayload{ID: team.ID, Name: team.Name})
								}
							}
							accessTokenPayload["teams"] = userTeams
							accessTokenPayload["userId"] = userID
							accessTokenPayload["isOnboarded"] = user.IsOnboarded
							return originalCreateNewSession(userID, accessTokenPayload, sessionDataInDatabase, disableAntiCsrf, tenantId, userContext)
						}

						return originalImplementation

					},
				},
				// ExposeAccessTokenToFrontendInCookieBasedAuth: true,
				CookieSecure:   &cookieSecure,
				CookieDomain:   &cookieDomain,
				CookieSameSite: &cookieSameSite,
			}),
		},
	})

	if err != nil {
		panic(err.Error()) //TODO resolve this. Use Fiber recover
	}
}
