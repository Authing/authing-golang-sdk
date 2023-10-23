package dto


type GetUserBatchDto struct{
    UserIds []string `json:"userIds,omitempty"`
    UserIdType string `json:"userIdType,omitempty"`
    WithCustomData bool `json:"withCustomData,omitempty"`
    FlatCustomData bool `json:"flatCustomData,omitempty"`
    WithIdentities bool `json:"withIdentities,omitempty"`
    WithDepartmentIds bool `json:"withDepartmentIds,omitempty"`
}

