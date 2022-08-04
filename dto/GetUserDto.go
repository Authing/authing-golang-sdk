package dto

type GetUserDto struct {
	UserId            string `json:"userId,omitempty"`
	UserIdType        string `json:"userIdType,omitempty"`
	WithCustomData    bool   `json:"withCustomData,omitempty"`
	WithIdentities    bool   `json:"withIdentities,omitempty"`
	WithDepartmentIds bool   `json:"withDepartmentIds,omitempty"`
}
