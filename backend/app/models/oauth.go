package models

type OauthAuthorizationRequest struct {
	ResponseType        string `json:"response_type"`
	ClientID            string `json:"client_id"`
	RedirectUri         string `json:"redirect_uri"`
	Scope               string `json:"scope"`
	State               string `json:"state"`
	CodeChallenge       string `json:"code_challenge"`
	CodeChallengeMethod string `json:"code_challenge_method"`
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
