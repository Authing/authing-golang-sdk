package dto


type AuthenticationOptionsDto struct{
    AuthenticationOptions  PublicKeyCredentialRequestOptionsDto `json:"authenticationOptions"`
    Ticket  string `json:"ticket"`
}

