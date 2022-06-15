package dto

type UserDepartmentPagingDto struct {
	TotalCount int                     `json:"totalCount"`
	List       []UserDepartmentRespDto `json:"list"`
}
