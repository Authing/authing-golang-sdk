package dto


type RevokeRoleBatchDto struct{
    Targets  []TargetDto `json:"targets"`
    Roles  []RoleCodeDto `json:"roles"`
}

