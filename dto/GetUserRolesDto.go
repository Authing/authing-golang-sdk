package dto

type GetUserRolesDto struct {
	UserId    string `json:"userId,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}
