package dto

type ListUsersDto struct {
	Page              int  `json:"page,omitempty"`
	Limit             int  `json:"limit,omitempty"`
	WithCustomData    bool `json:"withCustomData,omitempty"`
	WithIdentities    bool `json:"withIdentities,omitempty"`
	WithDepartmentIds bool `json:"withDepartmentIds,omitempty"`
}
