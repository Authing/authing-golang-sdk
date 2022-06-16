package dto

type RoleDepartmentListPagingDto struct {
	TotalCount int                     `json:"totalCount"`
	List       []RoleDepartmentRespDto `json:"list"`
}
