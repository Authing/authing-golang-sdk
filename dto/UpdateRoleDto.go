package dto


type UpdateRoleDto struct{
    NewCode  string `json:"newCode"`
    Code  string `json:"code"`
    Namespace  string `json:"namespace,omitempty"`
    Description  string `json:"description,omitempty"`
}

