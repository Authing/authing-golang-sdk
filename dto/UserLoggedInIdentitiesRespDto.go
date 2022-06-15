package dto

type UserLoggedInIdentitiesRespDto struct {
	StatusCode int                         `json:"statusCode"`
	Message    string                      `json:"message"`
	ApiCode    int                         `json:"apiCode,omitempty"`
	Data       []UserLoggedInIdentitiesDto `json:"data"`
}
