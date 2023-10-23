package dto


type RolePermissionListDto struct{
    RoleId  string `json:"roleId"`
    Status  string `json:"status"`
    EnableTime  int `json:"enableTime,omitempty"`
    EndTime  int `json:"endTime,omitempty"`
    UserPoolId  string `json:"userPoolId"`
    RoleName  string `json:"roleName"`
    RoleCode  string `json:"roleCode"`
    Description  string `json:"description,omitempty"`
    CreatedAt  string `json:"createdAt"`
    UpdatedAt  string `json:"updatedAt"`
}

