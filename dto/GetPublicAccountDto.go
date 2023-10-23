package dto


type GetPublicAccountDto struct{
    UserId string `json:"userId,omitempty"`
    UserIdType string `json:"userIdType,omitempty"`
    WithCustomData bool `json:"withCustomData,omitempty"`
    WithDepartmentIds bool `json:"withDepartmentIds,omitempty"`
}

