package dto


type PublicKeyCredentialDescriptorDto struct{
    Id  string `json:"id"`
    Transports  []string `json:"transports"`
    Type  string `json:"type"`
}

