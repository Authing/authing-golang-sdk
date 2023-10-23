package dto


type SetUserRolesDto struct{
    RoleIds  []string `json:"roleIds"`
    UserId  string `json:"userId"`
}

