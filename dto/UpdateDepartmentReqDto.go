package dto

type UpdateDepartmentReqDto struct {
	OrganizationCode   string `json:"organizationCode"`
	DepartmentId       string `json:"departmentId"`
	Description        string `json:"description,omitempty"`
	Code               string `json:"code,omitempty"`
	LeaderUserId       string `json:"leaderUserId,omitempty"`
	Name               string `json:"name,omitempty"`
	DepartmentIdType   string `json:"departmentIdType,omitempty"`
	ParentDepartmentId string `json:"parentDepartmentId,omitempty"`
}
