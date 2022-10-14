package dto


type UpdatePasswordDto struct{
    NewPassword  string `json:"newPassword"`
    OldPassword  string `json:"oldPassword,omitempty"`
    PasswordEncryptType  string  `json:"passwordEncryptType,omitempty"`
}

