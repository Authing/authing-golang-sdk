package dto


type GetProfileDto struct{
    WithCustomData bool `json:"withCustomData,omitempty"`
    WithIdentities bool `json:"withIdentities,omitempty"`
    WithDepartmentIds bool `json:"withDepartmentIds,omitempty"`
}

