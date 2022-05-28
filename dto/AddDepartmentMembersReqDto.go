package dto


type AddDepartmentMembersReqDto struct{
    UserIds  []string `json:"userIds"`
    OrganizationCode  string `json:"organizationCode"`
    DepartmentId  string `json:"departmentId"`
    DepartmentIdType  string  `json:"departmentIdType,omitempty"`
}

