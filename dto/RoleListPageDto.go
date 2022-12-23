package dto


type RoleListPageDto struct{
    TotalCount  int `json:"totalCount,omitempty"`
    Data  []RolePermissionListDto `json:"data"`
}

