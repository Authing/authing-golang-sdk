package dto


type RoleListItem struct{
    Code  string `json:"code"`
    Name  string `json:"name,omitempty"`
    Description  string `json:"description,omitempty"`
    Namespace  string `json:"namespace,omitempty"`
}

