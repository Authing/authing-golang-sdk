package dto

type ExtIdpSingleRespDto struct {
	StatusCode int       `json:"statusCode"`
	Message    string    `json:"message"`
	ApiCode    int       `json:"apiCode,omitempty"`
	Data       ExtIdpDto `json:"data"`
}
