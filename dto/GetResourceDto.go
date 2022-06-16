package dto

type GetResourceDto struct {
	Code      string `json:"code,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}
