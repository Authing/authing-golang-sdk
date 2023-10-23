package dto


type VerifyRegistrationDto struct{
    Ticket  string `json:"ticket"`
    RegistrationCredential  RegistrationCredentialDto `json:"registrationCredential"`
    AuthenticatorCode  string  `json:"authenticatorCode,omitempty"`
}

