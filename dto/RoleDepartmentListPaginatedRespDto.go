package dto

type RoleDepartmentListPaginatedRespDto struct {
	StatusCode int                         `json:"statusCode"`
	Message    string                      `json:"message"`
	ApiCode    int                         `json:"apiCode,omitempty"`
	Data       RoleDepartmentListPagingDto `json:"data"`
}
