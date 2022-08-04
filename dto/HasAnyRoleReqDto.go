package dto

type HasAnyRoleReqDto struct {
	Roles   []HasRoleRolesDto    `json:"roles"`
	UserId  string               `json:"userId"`
	Options HasAnyRoleOptionsDto `json:"options,omitempty"`
}
