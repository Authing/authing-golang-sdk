package dto

type GetUserDto struct {
	UserId            string `json:"userId,omitempty"`
	WithCustomData    bool   `json:"withCustomData,omitempty"`
	WithIdentities    bool   `json:"withIdentities,omitempty"`
	WithDepartmentIds bool   `json:"withDepartmentIds,omitempty"`
	Phone             string `json:"phone,omitempty"`
	Email             string `json:"email,omitempty"`
	Username          string `json:"username,omitempty"`
	ExternalId        string `json:"externalId,omitempty"`
}
