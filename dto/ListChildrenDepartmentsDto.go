package dto

type ListChildrenDepartmentsDto struct {
	OrganizationCode   string `json:"organizationCode,omitempty"`
	DepartmentId       string `json:"departmentId,omitempty"`
	DepartmentIdType   string `json:"departmentIdType,omitempty"`
	ExcludeVirtualNode bool   `json:"excludeVirtualNode,omitempty"`
	OnlyVirtualNode    bool   `json:"onlyVirtualNode,omitempty"`
	WithCustomData     bool   `json:"withCustomData,omitempty"`
	TenantId           string `json:"tenantId,omitempty"`
}
