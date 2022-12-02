package dto


type CreatePermissionNamespaceDto struct{
    Name  string `json:"name"`
    Code  string `json:"code"`
    Description  string `json:"description,omitempty"`
}

