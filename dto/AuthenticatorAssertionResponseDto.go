package dto


type AuthenticatorAssertionResponseDto struct{
    AuthenticatorData  string `json:"authenticatorData"`
    ClientDataJSON  string `json:"clientDataJSON"`
    Signature  string `json:"signature"`
    UserHandle  string `json:"userHandle"`
}

