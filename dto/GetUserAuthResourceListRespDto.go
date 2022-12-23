package dto

type GetUserAuthResourceListRespDto struct {
	StatusCode int                            `json:"statusCode"`
	Message    string                         `json:"message"`
	ApiCode    int                            `json:"apiCode,omitempty"`
	Data       GetUserAuthResourceListDataDto `json:"data"`
}
