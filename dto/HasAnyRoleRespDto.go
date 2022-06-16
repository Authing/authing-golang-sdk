package dto

type HasAnyRoleRespDto struct {
	StatusCode int           `json:"statusCode"`
	Message    string        `json:"message"`
	ApiCode    int           `json:"apiCode,omitempty"`
	Data       HasAnyRoleDto `json:"data"`
}
