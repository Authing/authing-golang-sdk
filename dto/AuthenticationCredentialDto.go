package dto


type AuthenticationCredentialDto struct{
    Id  string `json:"id"`
    RawId  string `json:"rawId"`
    Response  AuthenticatorAssertionResponseDto `json:"response"`
    Type  string `json:"type"`
}

