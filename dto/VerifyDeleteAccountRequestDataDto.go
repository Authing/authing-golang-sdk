package dto


type VerifyDeleteAccountRequestDataDto struct{
    DeleteAccountToken  string `json:"deleteAccountToken"`
    TokenExpiresIn  int `json:"tokenExpiresIn"`
}

