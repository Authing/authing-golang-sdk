package dto


type RegistrationOptionsDto struct{
    RegistrationOptions  PublicKeyCredentialCreationOptionsDto `json:"registrationOptions"`
    Ticket  string `json:"ticket"`
}

