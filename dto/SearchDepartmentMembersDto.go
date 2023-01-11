package dto

type SearchDepartmentMembersDto struct {
	OrganizationCode           string `json:"organizationCode,omitempty"`
	DepartmentId               string `json:"departmentId,omitempty"`
	Keywords                   string `json:"keywords,omitempty"`
	Page                       int    `json:"page,omitempty"`
	Limit                      int    `json:"limit,omitempty"`
	DepartmentIdType           string `json:"departmentIdType,omitempty"`
	IncludeChildrenDepartments bool   `json:"includeChildrenDepartments,omitempty"`
	WithCustomData             bool   `json:"withCustomData,omitempty"`
	WithIdentities             bool   `json:"withIdentities,omitempty"`
	WithDepartmentIds          bool   `json:"withDepartmentIds,omitempty"`
	TenantId                   string `json:"tenantId,omitempty"`
}
