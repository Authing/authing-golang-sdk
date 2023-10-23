package dto


type RoleDto struct{
    Id  string `json:"id"`
    Code  string `json:"code"`
    Name  string `json:"name"`
    Description  string `json:"description"`
    Namespace  string `json:"namespace"`
    NamespaceName  string `json:"namespaceName"`
    Status  string `json:"status,omitempty"`
    DisableTime  int `json:"disableTime,omitempty"`
}

