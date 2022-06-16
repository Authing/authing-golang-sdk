package dto

type UserMfaSingleRespDto struct {
	StatusCode int            `json:"statusCode"`
	Message    string         `json:"message"`
	ApiCode    int            `json:"apiCode,omitempty"`
	Data       UserMfaRespDto `json:"data"`
}
