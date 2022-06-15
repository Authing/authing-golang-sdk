package dto

type SetUserCustomDataRespDto struct {
	StatusCode int                  `json:"statusCode"`
	Message    string               `json:"message"`
	ApiCode    int                  `json:"apiCode,omitempty"`
	Data       SetUserCustomDataDto `json:"data"`
}
