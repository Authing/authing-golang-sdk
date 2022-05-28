package dto


type ListChildrenDepartmentsDto struct{
    DepartmentId string `json:"department_id,omitempty"`
    OrganizationCode string `json:"organization_code,omitempty"`
    DepartmentIdType string `json:"department_id_type,omitempty"`
}

