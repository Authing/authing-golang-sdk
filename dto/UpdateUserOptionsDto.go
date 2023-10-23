package dto


type UpdateUserOptionsDto struct{
    UserIdType  string  `json:"userIdType,omitempty"`
    ResetPasswordOnFirstLogin  bool `json:"resetPasswordOnFirstLogin,omitempty"`
    ResetPasswordOnNextLogin  bool `json:"resetPasswordOnNextLogin,omitempty"`
    PasswordEncryptType  string  `json:"passwordEncryptType,omitempty"`
    AutoGeneratePassword  bool `json:"autoGeneratePassword,omitempty"`
    SendPasswordResetedNotification  SendResetPasswordNotificationDto `json:"sendPasswordResetedNotification,omitempty"`
}

