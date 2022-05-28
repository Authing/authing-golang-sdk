package dto

type UsersListPagingDto struct {
	TotalCount int       `json:"totalCount"`
	List       []UserDto `json:"list"`
}
