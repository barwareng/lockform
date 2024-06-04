package models

type OauthAuthorizationRequest struct {
	ResponseType        string `json:"response_type" query:"response_type"`
	ClientID            string `json:"client_id" query:"client_id"`
	RedirectUri         string `json:"redirect_uri" query:"redirect_uri"`
	Scope               string
	State               string
	CodeChallenge       string `json:"code_challenge" query:"code_challenge"`
	CodeChallengeMethod string `json:"code_challenge_method" query:"code_challenge_method"`
}
type OauthAuthorizationResponse struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

type OauthTokenRequest struct {
	GrantType    string `json:"grant_type" validate:"required"`
	Code         string `json:"code" validate:"required"`
	RedirectUri  string `json:"redirect_uri" validate:"required"`
	CodeVerifier string `json:"code_verifier"`
}

type OauthTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
}
