package dto

type DeleteResourceDto struct {
	Code      string `json:"code"`
	Namespace string `json:"namespace,omitempty"`
}
