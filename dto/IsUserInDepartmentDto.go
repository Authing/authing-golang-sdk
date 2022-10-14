package dto


type IsUserInDepartmentDto struct{
    UserId string `json:"userId,omitempty"`
    OrganizationCode string `json:"organizationCode,omitempty"`
    DepartmentId string `json:"departmentId,omitempty"`
    DepartmentIdType string `json:"departmentIdType,omitempty"`
    IncludeChildrenDepartments bool `json:"includeChildrenDepartments,omitempty"`
}

