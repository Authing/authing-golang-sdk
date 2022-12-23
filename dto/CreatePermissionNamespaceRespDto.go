package dto


type CreatePermissionNamespaceRespDto struct{
    Code  string `json:"code"`
    Name  string `json:"name"`
    Description  string `json:"description,omitempty"`
}

