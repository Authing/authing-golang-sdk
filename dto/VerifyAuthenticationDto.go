package dto


type VerifyAuthenticationDto struct{
    Ticket  string `json:"ticket"`
    AuthenticationCredential  AuthenticationCredentialDto `json:"authenticationCredential"`
    Options  SignInByWebAuthnOptionsDto `json:"options,omitempty"`
}

