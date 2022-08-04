package dto

type GetUserMfaInfoDto struct {
	UserId     string `json:"userId,omitempty"`
	UserIdType string `json:"userIdType,omitempty"`
}
