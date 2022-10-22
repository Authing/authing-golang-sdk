package dto


type TokenEndPointParams struct{
    ClientId  string `json:"client_id,omitempty"`
    ClientSecret  string `json:"client_secret,omitempty"`
    GrantType  string  `json:"grant_type"`
    RedirectUri  string `json:"redirect_uri"`
    Code  string `json:"code,omitempty"`
    CodeVerifier  string `json:"code_verifier,omitempty"`
    RefreshToken  string `json:"refresh_token,omitempty"`
}

