package dto

type DeleteDepartmentReqDto struct {
	DepartmentId     string `json:"departmentId"`
	OrganizationCode string `json:"organizationCode"`
	DepartmentIdType string `json:"departmentIdType,omitempty"`
}
