package dto


type RegistrationCredentialDto struct{
    Id  string `json:"id"`
    RawId  string `json:"rawId"`
    Response  AuthenticatorAttestationResponseDto `json:"response"`
    Transports  []string `json:"transports,omitempty"`
    Type  string `json:"type"`
}

