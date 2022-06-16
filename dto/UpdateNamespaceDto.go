package dto

type UpdateNamespaceDto struct {
	Code        string `json:"code"`
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	NewCode     string `json:"newCode,omitempty"`
}
