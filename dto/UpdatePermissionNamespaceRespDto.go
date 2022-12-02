package dto


type UpdatePermissionNamespaceRespDto struct{
    Code  string `json:"code"`
    Name  string `json:"name"`
    Description  string `json:"description,omitempty"`
}

