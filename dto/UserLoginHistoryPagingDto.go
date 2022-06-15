package dto

type UserLoginHistoryPagingDto struct {
	TotalCount int                   `json:"totalCount"`
	List       []UserLoginHistoryDto `json:"list"`
}
