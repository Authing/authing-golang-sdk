package dto


type VerifyUpdateEmailRequestData struct{
    UpdateEmailToken  string `json:"updateEmailToken"`
    TokenExpiresIn  int `json:"tokenExpiresIn"`
}

