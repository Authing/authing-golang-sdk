package dto


type CreateAdminRoleDto struct{
    Name  string `json:"name"`
    Code  string `json:"code"`
    Description  string `json:"description,omitempty"`
}

