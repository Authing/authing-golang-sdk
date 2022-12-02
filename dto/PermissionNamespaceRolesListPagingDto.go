package dto


type PermissionNamespaceRolesListPagingDto struct{
    TotalCount  int `json:"totalCount"`
    List  []PermissionNamespaceRolesListRespDto `json:"list"`
}

