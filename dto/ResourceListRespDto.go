package dto

type ResourceListRespDto struct {
	StatusCode int           `json:"statusCode"`
	Message    string        `json:"message"`
	ApiCode    int           `json:"apiCode,omitempty"`
	Data       []ResourceDto `json:"data"`
}
