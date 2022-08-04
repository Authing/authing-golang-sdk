package dto

type IsActionAllowedRespDtp struct {
	StatusCode int                    `json:"statusCode"`
	Message    string                 `json:"message"`
	ApiCode    int                    `json:"apiCode,omitempty"`
	Data       IsActionAllowedDataDto `json:"data"`
}
