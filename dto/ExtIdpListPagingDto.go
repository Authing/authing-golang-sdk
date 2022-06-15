package dto

type ExtIdpListPagingDto struct {
	TotalCount int         `json:"totalCount"`
	List       []ExtIdpDto `json:"list"`
}
