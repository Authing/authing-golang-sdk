package dto


type DeleteRoleBatchDto struct{
    RoleList  []RoleCodeAndNamespaceDto `json:"roleList"`
}

