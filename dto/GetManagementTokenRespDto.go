package dto

type GetManagementTokenRespDto struct {
	StatusCode int            `json:"statusCode"`
	Message    string         `json:"message"`
	ApiCode    int            `json:"apiCode,omitempty"`
	Data       AccessTokenDto `json:"data"`
}
