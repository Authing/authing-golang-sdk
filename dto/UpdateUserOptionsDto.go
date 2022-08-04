package dto

type UpdateUserOptionsDto struct {
	UserIdType                      string                           `json:"userIdType,omitempty"`
	ResetPasswordOnNextLogin        bool                             `json:"resetPasswordOnNextLogin,omitempty"`
	AutoGeneratePassword            bool                             `json:"autoGeneratePassword,omitempty"`
	SendPasswordResetedNotification SendResetPasswordNotificationDto `json:"sendPasswordResetedNotification,omitempty"`
}
