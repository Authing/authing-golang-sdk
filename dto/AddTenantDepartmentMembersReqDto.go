package dto


type AddTenantDepartmentMembersReqDto struct{
    OrganizationCode  string `json:"organizationCode"`
    DepartmentId  string `json:"departmentId"`
    DepartmentIdType  string  `json:"departmentIdType,omitempty"`
    LinkUserIds  []string `json:"linkUserIds,omitempty"`
    MemberIds  []string `json:"memberIds,omitempty"`
    TenantId  string `json:"tenantId,omitempty"`
}

