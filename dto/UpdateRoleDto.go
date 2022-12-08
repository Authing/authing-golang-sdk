package dto


type UpdateRoleDto struct{
    Name  string `json:"name"`
    NewCode  string `json:"newCode"`
    Code  string `json:"code"`
    Namespace  string `json:"namespace,omitempty"`
    Description  string `json:"description,omitempty"`
    Status  string `json:"status,omitempty"`
}

