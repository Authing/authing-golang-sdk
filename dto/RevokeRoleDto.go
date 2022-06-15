package dto

type RevokeRoleDto struct {
	Targets   []TargetDto `json:"targets"`
	Code      string      `json:"code"`
	Namespace string      `json:"namespace,omitempty"`
}
