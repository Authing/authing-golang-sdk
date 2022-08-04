package dto

type IsUserInDepartmentRespDto struct {
	StatusCode int                       `json:"statusCode"`
	Message    string                    `json:"message"`
	ApiCode    int                       `json:"apiCode,omitempty"`
	Data       IsUserInDepartmentDataDto `json:"data"`
}
