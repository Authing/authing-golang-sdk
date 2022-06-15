package dto

type GetCustomDataRespDto struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	ApiCode    int         `json:"apiCode,omitempty"`
	Data       interface{} `json:"data"`
}
