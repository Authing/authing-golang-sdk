package dto

type UserMfaRespDto struct {
	TotpStatus    string `json:"totpStatus"`
	FaceMfaStatus string `json:"faceMfaStatus"`
}
