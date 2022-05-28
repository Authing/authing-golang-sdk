package dto


type GetUserBatchDto struct{
    UserIds string `json:"user_ids,omitempty"`
    WithCustomData bool `json:"with_custom_data,omitempty"`
    WithIdentities bool `json:"with_identities,omitempty"`
    WithDepartmentIds bool `json:"with_department_ids,omitempty"`
}

