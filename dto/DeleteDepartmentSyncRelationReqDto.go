package dto


type DeleteDepartmentSyncRelationReqDto struct{
    Provider  string `json:"provider"`
    DepartmentId  string `json:"departmentId"`
    OrganizationCode  string `json:"organizationCode"`
    DepartmentIdType  string  `json:"departmentIdType,omitempty"`
}

