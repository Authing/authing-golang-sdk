package dto

type CheckAppIdSecretDataDto struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message,omitempty"`
}
