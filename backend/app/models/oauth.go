package models

type OathAuthorizationRequest struct {
	ResponseType        string `json:"response_type" validate:"required"`
	ClientID            string `json:"client_id" validate:"required"`
	RedirectUri         string `json:"redirect_uri" validate:"required"`
	Scope               string `json:"scope"`
	State               string `json:"state"`
	CodeChallenge       string `json:"code_challenge"`
	CodeChallengeMethod string `json:"code_challenge_method"`
}
type OathAuthorizationResponse struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

// GET {Authorization Endpoint}
//   ?response_type=code             // - Required
//   &client_id={Client ID}          // - Required
//   &redirect_uri={Redirect URI}    // - Conditionally required
//   &scope={Scopes}                 // - Optional
//   &state={Arbitrary String}       // - Recommended
//   &code_challenge={Challenge}     // - Optional
//   &code_challenge_method={Method} // - Optional
//   HTTP/1.1
// HOST: {Authorization Server}
