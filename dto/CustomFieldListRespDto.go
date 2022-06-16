package dto

type CustomFieldListRespDto struct {
	StatusCode int              `json:"statusCode"`
	Message    string           `json:"message"`
	ApiCode    int              `json:"apiCode,omitempty"`
	Data       []CustomFieldDto `json:"data"`
}
