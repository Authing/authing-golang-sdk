package dto


type SendResetPasswordNotificationDto struct{
    SendDefaultEmailNotification  bool `json:"sendDefaultEmailNotification,omitempty"`
    SendDefaultPhoneNotification  bool `json:"sendDefaultPhoneNotification,omitempty"`
    InputSendEmailNotification  string `json:"inputSendEmailNotification,omitempty"`
    InputSendPhoneNotification  string `json:"inputSendPhoneNotification,omitempty"`
    AppId  string `json:"appId,omitempty"`
}

