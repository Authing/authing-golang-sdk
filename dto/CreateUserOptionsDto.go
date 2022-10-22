package dto


type CreateUserOptionsDto struct{
    KeepPassword  bool `json:"keepPassword,omitempty"`
    AutoGeneratePassword  bool `json:"autoGeneratePassword,omitempty"`
    ResetPasswordOnFirstLogin  bool `json:"resetPasswordOnFirstLogin,omitempty"`
    DepartmentIdType  string  `json:"departmentIdType,omitempty"`
    SendNotification  SendCreateAccountNotificationDto `json:"sendNotification,omitempty"`
    PasswordEncryptType  string  `json:"passwordEncryptType,omitempty"`
}

