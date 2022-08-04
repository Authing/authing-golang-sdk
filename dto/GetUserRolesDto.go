package dto

type GetUserRolesDto struct {
	UserId     string `json:"userId,omitempty"`
	UserIdType string `json:"userIdType,omitempty"`
	Namespace  string `json:"namespace,omitempty"`
}
