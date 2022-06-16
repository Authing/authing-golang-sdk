package dto

type CreateRoleDto struct {
	Code        string `json:"code"`
	Namespace   string `json:"namespace,omitempty"`
	Description string `json:"description,omitempty"`
}
