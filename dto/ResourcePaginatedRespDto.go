package dto

type ResourcePaginatedRespDto struct {
	StatusCode int               `json:"statusCode"`
	Message    string            `json:"message"`
	ApiCode    int               `json:"apiCode,omitempty"`
	Data       ResourcePagingDto `json:"data"`
}
