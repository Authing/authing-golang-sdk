package dto

type GroupPagingDto struct {
	TotalCount int           `json:"totalCount"`
	List       []ResGroupDto `json:"list"`
}
