package dto

type GetUserAccessibleAppsDto struct {
	UserId     string `json:"userId,omitempty"`
	UserIdType string `json:"userIdType,omitempty"`
}
