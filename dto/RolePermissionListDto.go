package dto


type RolePermissionListDto struct{
    RoleId  string `json:"roleId"`
    UserPoolId  string `json:"userPoolId"`
    RoleName  string `json:"roleName"`
    RoleCode  string `json:"roleCode"`
    Description  string `json:"description,omitempty"`
    AppId  string `json:"appId"`
    AppName  string `json:"appName"`
    CreatedAt  string `json:"createdAt"`
    UpdatedAt  string `json:"updatedAt"`
}

