package dto

type UserDepartmentPaginatedRespDto struct {
	StatusCode int                     `json:"statusCode"`
	Message    string                  `json:"message"`
	ApiCode    int                     `json:"apiCode,omitempty"`
	Data       UserDepartmentPagingDto `json:"data"`
}
