package dto

type ListDepartmentMembersDto struct {
	OrganizationCode           string `json:"organizationCode,omitempty"`
	DepartmentId               string `json:"departmentId,omitempty"`
	SortBy                     string `json:"sortBy,omitempty"`
	OrderBy                    string `json:"orderBy,omitempty"`
	DepartmentIdType           string `json:"departmentIdType,omitempty"`
	IncludeChildrenDepartments bool   `json:"includeChildrenDepartments,omitempty"`
	Page                       int    `json:"page,omitempty"`
	Limit                      int    `json:"limit,omitempty"`
	WithCustomData             bool   `json:"withCustomData,omitempty"`
	WithIdentities             bool   `json:"withIdentities,omitempty"`
	WithDepartmentIds          bool   `json:"withDepartmentIds,omitempty"`
	TenantId                   string `json:"tenantId,omitempty"`
}
