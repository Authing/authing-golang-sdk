package dto


type CheckRoleParamsDto struct{
    Code  string `json:"code"`
    Namespace  string `json:"namespace"`
    Name  string `json:"name,omitempty"`
}

