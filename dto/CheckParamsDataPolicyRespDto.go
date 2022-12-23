package dto

type CheckParamsDataPolicyRespDto struct {
	IsValid bool   `json:"isValid"`
	Message string `json:"message,omitempty"`
}
