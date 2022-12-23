package dto


type PermissionNamespaceListPagingDto struct{
    TotalCount  int `json:"totalCount"`
    List  []PermissionNamespacesListRespDto `json:"list"`
}

