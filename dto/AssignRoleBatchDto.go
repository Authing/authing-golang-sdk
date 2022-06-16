package dto

type AssignRoleBatchDto struct {
	Targets []TargetDto   `json:"targets"`
	Roles   []RoleCodeDto `json:"roles"`
}
