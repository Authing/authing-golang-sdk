package dto

type GroupPaginatedRespDto struct {
	StatusCode int            `json:"statusCode"`
	Message    string         `json:"message"`
	ApiCode    int            `json:"apiCode,omitempty"`
	Data       GroupPagingDto `json:"data"`
}
