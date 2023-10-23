package dto


type GetUserDto struct{
    UserId string `json:"userId,omitempty"`
    UserIdType string `json:"userIdType,omitempty"`
    FlatCustomData bool `json:"flatCustomData,omitempty"`
    WithCustomData bool `json:"withCustomData,omitempty"`
    WithPost bool `json:"withPost,omitempty"`
    WithIdentities bool `json:"withIdentities,omitempty"`
    WithDepartmentIds bool `json:"withDepartmentIds,omitempty"`
}

