package dto


type ListDepartmentMembersDto struct{
    OrganizationCode string `json:"organization_code,omitempty"`
    DepartmentId string `json:"department_id,omitempty"`
    DepartmentIdType string `json:"department_id_type,omitempty"`
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
    WithCustomData bool `json:"with_custom_data,omitempty"`
    WithIdentities bool `json:"with_identities,omitempty"`
    WithDepartmentIds bool `json:"with_department_ids,omitempty"`
}

