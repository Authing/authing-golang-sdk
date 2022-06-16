package dto

type IsSuccessRespDto struct {
	StatusCode int          `json:"statusCode"`
	Message    string       `json:"message"`
	ApiCode    int          `json:"apiCode,omitempty"`
	Data       IsSuccessDto `json:"data"`
}
