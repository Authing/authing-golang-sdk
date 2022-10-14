package dto


type VerifyUpdatePhoneRequestData struct{
    UpdatePhoneToken  string `json:"updatePhoneToken"`
    TokenExpiresIn  int `json:"tokenExpiresIn"`
}

