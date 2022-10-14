package dto


type ResetPasswordVerify struct{
    PasswordResetToken  string `json:"passwordResetToken"`
    TokenExpiresIn  int `json:"tokenExpiresIn"`
}

