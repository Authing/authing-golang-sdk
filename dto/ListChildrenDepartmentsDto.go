package dto

type ListChildrenDepartmentsDto struct {
	DepartmentId     string `json:"departmentId,omitempty"`
	OrganizationCode string `json:"organizationCode,omitempty"`
	DepartmentIdType string `json:"departmentIdType,omitempty"`
}
