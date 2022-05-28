package dto


type CreateDepartmentReqDto struct{
    OrganizationCode  string `json:"organizationCode"`
    Name  string `json:"name"`
    ParentDepartmentId  string `json:"parentDepartmentId"`
    OpenDepartmentId  string `json:"openDepartmentId,omitempty"`
    Code  string `json:"code,omitempty"`
    LeaderUserId  string `json:"leaderUserId,omitempty"`
    DepartmentIdType  string  `json:"departmentIdType,omitempty"`
}

