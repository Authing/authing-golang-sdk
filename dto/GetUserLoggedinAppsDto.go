package dto

type GetUserLoggedinAppsDto struct {
	UserId     string `json:"userId,omitempty"`
	UserIdType string `json:"userIdType,omitempty"`
}
