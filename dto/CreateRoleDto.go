package dto


type CreateRoleDto struct{
    Code  string `json:"code"`
    Name  string `json:"name,omitempty"`
    Namespace  string `json:"namespace,omitempty"`
    Description  string `json:"description,omitempty"`
    DisableTime  string `json:"disableTime,omitempty"`
}

