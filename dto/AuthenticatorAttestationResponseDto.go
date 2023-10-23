package dto


type AuthenticatorAttestationResponseDto struct{
    AttestationObject  string `json:"attestationObject"`
    ClientDataJSON  string `json:"clientDataJSON"`
}

