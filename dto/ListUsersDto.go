package dto


type ListUsersDto struct{
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
    Status string `json:"status,omitempty"`
    UpdatedAtStart int `json:"updatedAtStart,omitempty"`
    UpdatedAtEnd int `json:"updatedAtEnd,omitempty"`
    WithCustomData bool `json:"withCustomData,omitempty"`
    WithPost bool `json:"withPost,omitempty"`
    WithIdentities bool `json:"withIdentities,omitempty"`
    WithDepartmentIds bool `json:"withDepartmentIds,omitempty"`
}

