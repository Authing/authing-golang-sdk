package dto


type GetPublicAccountBatchDto struct{
    UserIds []string `json:"userIds,omitempty"`
    UserIdType string `json:"userIdType,omitempty"`
    WithCustomData bool `json:"withCustomData,omitempty"`
    WithDepartmentIds bool `json:"withDepartmentIds,omitempty"`
}

