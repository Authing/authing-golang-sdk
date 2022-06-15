package dto

type RoleListItem struct {
	Code        string `json:"code"`
	Description string `json:"description,omitempty"`
	Namespace   string `json:"namespace,omitempty"`
}
