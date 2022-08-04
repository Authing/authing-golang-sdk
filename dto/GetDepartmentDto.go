package dto

type GetDepartmentDto struct {
	OrganizationCode string `json:"organizationCode,omitempty"`
	DepartmentId     string `json:"departmentId,omitempty"`
	DepartmentCode   string `json:"departmentCode,omitempty"`
	DepartmentIdType string `json:"departmentIdType,omitempty"`
	WithCustomData   bool   `json:"withCustomData,omitempty"`
}
