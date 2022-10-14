package dto


type GetUserBatchDto struct{
    UserIds string `json:"userIds,omitempty"`
    UserIdType string `json:"userIdType,omitempty"`
    WithCustomData bool `json:"withCustomData,omitempty"`
    WithIdentities bool `json:"withIdentities,omitempty"`
    WithDepartmentIds bool `json:"withDepartmentIds,omitempty"`
}

