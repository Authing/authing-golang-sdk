package dto


type GetParentDepartmentDto struct{
    OrganizationCode string `json:"organization_code,omitempty"`
    DepartmentId string `json:"department_id,omitempty"`
    DepartmentIdType string `json:"department_id_type,omitempty"`
}

