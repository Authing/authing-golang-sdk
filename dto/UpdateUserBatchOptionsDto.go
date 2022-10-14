package dto


type UpdateUserBatchOptionsDto struct{
    ResetPasswordOnNextLogin  bool `json:"resetPasswordOnNextLogin,omitempty"`
    PasswordEncryptType  string  `json:"passwordEncryptType,omitempty"`
    AutoGeneratePassword  bool `json:"autoGeneratePassword,omitempty"`
    SendPasswordResetedNotification  SendResetPasswordNotificationDto `json:"sendPasswordResetedNotification,omitempty"`
}

