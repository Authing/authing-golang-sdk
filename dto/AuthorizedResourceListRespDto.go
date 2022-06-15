package dto

type AuthorizedResourceListRespDto struct {
	StatusCode int                     `json:"statusCode"`
	Message    string                  `json:"message"`
	ApiCode    int                     `json:"apiCode,omitempty"`
	Data       []AuthorizedResourceDto `json:"data"`
}
