package dto

type UserPaginatedRespDto struct {
	StatusCode int           `json:"statusCode"`
	Message    string        `json:"message"`
	ApiCode    int           `json:"apiCode,omitempty"`
	Data       UserPagingDto `json:"data"`
}
