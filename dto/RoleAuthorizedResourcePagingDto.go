package dto

type RoleAuthorizedResourcePagingDto struct {
	TotalCount int                              `json:"totalCount"`
	List       []RoleAuthorizedResourcesRespDto `json:"list"`
}
