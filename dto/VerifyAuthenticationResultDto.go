package dto


type VerifyAuthenticationResultDto struct{
    Verified  bool `json:"verified"`
    TokenSet  LoginTokenResponseDataDto `json:"tokenSet,omitempty"`
}

