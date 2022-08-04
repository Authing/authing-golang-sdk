package dto

type GetParentDepartmentDto struct {
	OrganizationCode string `json:"organizationCode,omitempty"`
	DepartmentId     string `json:"departmentId,omitempty"`
	DepartmentIdType string `json:"departmentIdType,omitempty"`
	WithCustomData   bool   `json:"withCustomData,omitempty"`
}
