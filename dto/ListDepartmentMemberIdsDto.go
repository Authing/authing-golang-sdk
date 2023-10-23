package dto


type ListDepartmentMemberIdsDto struct{
    OrganizationCode string `json:"organizationCode,omitempty"`
    DepartmentId string `json:"departmentId,omitempty"`
    DepartmentIdType string `json:"departmentIdType,omitempty"`
    TenantId string `json:"tenantId,omitempty"`
}

