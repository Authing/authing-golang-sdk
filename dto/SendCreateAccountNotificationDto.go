package dto

type SendCreateAccountNotificationDto struct {
	SendEmailNotification bool   `json:"sendEmailNotification,omitempty"`
	SendPhoneNotification bool   `json:"sendPhoneNotification,omitempty"`
	AppId                 string `json:"appId,omitempty"`
}
