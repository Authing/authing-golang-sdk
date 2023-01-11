package dto

type CheckUserSameLevelPermissionRespDto struct {
	Action           string `json:"action"`
	ResourceNodeCode string `json:"resourceNodeCode"`
	Enabled          bool   `json:"enabled"`
}
