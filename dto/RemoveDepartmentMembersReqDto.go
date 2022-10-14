package dto


type RemoveDepartmentMembersReqDto struct{
    UserIds  []string `json:"userIds"`
    OrganizationCode  string `json:"organizationCode"`
    DepartmentId  string `json:"departmentId"`
    DepartmentIdType  string  `json:"departmentIdType,omitempty"`
}

