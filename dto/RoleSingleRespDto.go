package dto

type RoleSingleRespDto struct {
	StatusCode int     `json:"statusCode"`
	Message    string  `json:"message"`
	ApiCode    int     `json:"apiCode,omitempty"`
	Data       RoleDto `json:"data"`
}
