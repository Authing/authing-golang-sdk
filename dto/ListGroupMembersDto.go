package dto


type ListGroupMembersDto struct{
    Code string `json:"code,omitempty"`
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
    WithCustomData bool `json:"with_custom_data,omitempty"`
    WithIdentities bool `json:"with_identities,omitempty"`
    WithDepartmentIds bool `json:"with_department_ids,omitempty"`
}

