package dto


type UpdateEmailByEmailPassCodeDto struct{
    NewEmail  string `json:"newEmail"`
    NewEmailPassCode  string `json:"newEmailPassCode"`
    OldEmail  string `json:"oldEmail,omitempty"`
    OldEmailPassCode  string `json:"oldEmailPassCode,omitempty"`
}

