package dto


type GetUserDto struct{
    UserId string `json:"user_id,omitempty"`
    WithCustomData bool `json:"with_custom_data,omitempty"`
    WithIdentities bool `json:"with_identities,omitempty"`
    WithDepartmentIds bool `json:"with_department_ids,omitempty"`
    Phone string `json:"phone,omitempty"`
    Email string `json:"email,omitempty"`
    Username string `json:"username,omitempty"`
    ExternalId string `json:"externalId,omitempty"`
}

