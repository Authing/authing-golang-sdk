package dto


type PermissionNamespaceRolesListRespDto struct{
    Name  string `json:"name"`
    Code  string `json:"code"`
    Description  string `json:"description,omitempty"`
    Namespace  string `json:"namespace"`
    UpdatedAt  string `json:"updatedAt"`
}

