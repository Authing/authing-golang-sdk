package dto

type DepartmentSingleRespDto struct {
	StatusCode int           `json:"statusCode"`
	Message    string        `json:"message"`
	ApiCode    int           `json:"apiCode,omitempty"`
	Data       DepartmentDto `json:"data"`
}
