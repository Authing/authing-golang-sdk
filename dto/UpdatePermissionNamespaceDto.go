package dto


type UpdatePermissionNamespaceDto struct{
    Code  string `json:"code"`
    Name  string `json:"name,omitempty"`
    NewCode  string `json:"newCode,omitempty"`
    Description  string `json:"description,omitempty"`
}

