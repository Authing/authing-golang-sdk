package dto

type IdentityListRespDto struct {
	StatusCode int           `json:"statusCode"`
	Message    string        `json:"message"`
	ApiCode    int           `json:"apiCode,omitempty"`
	Data       []IdentityDto `json:"data"`
}
