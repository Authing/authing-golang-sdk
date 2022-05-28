package dto

type UserListRespDto struct {
	StatusCode int                `json:"statusCode"`
	Message    string             `json:"message"`
	ApiCode    int                `json:"apiCode,omitempty"`
	Data       UsersListPagingDto `json:"data"`
}
