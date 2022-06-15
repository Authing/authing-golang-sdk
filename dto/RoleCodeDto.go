package dto

type RoleCodeDto struct {
	Code      string `json:"code"`
	Namespace string `json:"namespace,omitempty"`
}
