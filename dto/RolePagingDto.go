package dto

type RolePagingDto struct {
	TotalCount int       `json:"totalCount"`
	List       []RoleDto `json:"list"`
}
