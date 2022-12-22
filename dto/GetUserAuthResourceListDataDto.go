package dto

type GetUserAuthResourceListDataDto struct {
	UserPermissionList []UserAuthResourceListDto `json:"userPermissionList"`
}
