package dto


type UpdateDepartmentReqDto struct{
    OrganizationCode  string `json:"organizationCode"`
    ParentDepartmentId  string `json:"parentDepartmentId"`
    DepartmentId  string `json:"departmentId"`
    Code  string `json:"code,omitempty"`
    LeaderUserId  string `json:"leaderUserId,omitempty"`
    Name  string `json:"name,omitempty"`
    DepartmentIdType  string  `json:"departmentIdType,omitempty"`
}

