package dto


type CreateTenantUserOptionsDto struct{
    KeepPassword  bool `json:"keepPassword,omitempty"`
    AutoGeneratePassword  bool `json:"autoGeneratePassword,omitempty"`
    ResetPasswordOnFirstLogin  bool `json:"resetPasswordOnFirstLogin,omitempty"`
    PasswordEncryptType  string  `json:"passwordEncryptType,omitempty"`
}

