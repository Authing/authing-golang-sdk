package dto


type ResetPasswordDto struct{
    Password  string `json:"password"`
    PasswordResetToken  string `json:"passwordResetToken"`
    PasswordEncryptType  string  `json:"passwordEncryptType,omitempty"`
}

