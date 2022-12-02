package dto


type PermissionNamespacesListRespDto struct{
    Name  string `json:"name"`
    Code  string `json:"code"`
    Description  string `json:"description,omitempty"`
    Status  int `json:"status,omitempty"`
}

