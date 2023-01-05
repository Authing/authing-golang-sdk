package dto

type EnrollFactorEnrollmentDataDto struct {
	PassCode        string `json:"passCode,omitempty"`
	Photo           string `json:"photo,omitempty"`
	IsExternalPhoto bool   `json:"isExternalPhoto,omitempty"`
}
