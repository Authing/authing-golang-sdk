package dto

type SetCustomDataReqDto struct {
	List             []SetCustomDataDto `json:"list"`
	TargetIdentifier string             `json:"targetIdentifier"`
	TargetType       string             `json:"targetType"`
	Namespace        string             `json:"namespace,omitempty"`
}
