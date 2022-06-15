package dto

type NamespaceDto struct {
	Code        string `json:"code"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
