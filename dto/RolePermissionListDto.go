package dto


type RolePermissionListDto struct{
    RoleId  string `json:"roleId"`
    UserPoolId  string `json:"userPoolId"`
    RoleName  string `json:"roleName"`
    RoleCode  string `json:"roleCode"`
    Description  string `json:"description,omitempty"`
    CreatedAt  string `json:"createdAt"`
    UpdatedAt  string `json:"updatedAt"`
}

