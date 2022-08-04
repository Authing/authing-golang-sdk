package dto

type ListRolesDto struct {
	Keywords  string `json:"keywords,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Page      int    `json:"page,omitempty"`
	Limit     int    `json:"limit,omitempty"`
}
