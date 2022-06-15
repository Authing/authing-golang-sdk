package dto

type DeleteDepartmentReqDto struct {
	OrganizationCode string `json:"organizationCode"`
	DepartmentId     string `json:"departmentId"`
	DepartmentIdType string `json:"departmentIdType,omitempty"`
}
