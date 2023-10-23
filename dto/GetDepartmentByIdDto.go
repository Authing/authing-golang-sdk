package dto


type GetDepartmentByIdDto struct{
    DepartmentId string `json:"departmentId,omitempty"`
    TenantId string `json:"tenantId,omitempty"`
    WithCustomData bool `json:"withCustomData,omitempty"`
}

