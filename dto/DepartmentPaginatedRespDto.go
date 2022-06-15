package dto

type DepartmentPaginatedRespDto struct {
	StatusCode int                 `json:"statusCode"`
	Message    string              `json:"message"`
	ApiCode    int                 `json:"apiCode,omitempty"`
	Data       DepartmentPagingDto `json:"data"`
}
