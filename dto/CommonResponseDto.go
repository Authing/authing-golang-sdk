package dto

type CommonResponseDto struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	ApiCode    int    `json:"apiCode,omitempty"`
}
