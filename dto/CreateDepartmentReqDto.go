package dto

type CreateDepartmentReqDto struct {
	OrganizationCode   string `json:"organizationCode"`
	ParentDepartmentId string `json:"parentDepartmentId"`
	Name               string `json:"name"`
	OpenDepartmentId   string `json:"openDepartmentId,omitempty"`
	Description        string `json:"description,omitempty"`
	Code               string `json:"code,omitempty"`
	LeaderUserId       string `json:"leaderUserId,omitempty"`
	DepartmentIdType   string `json:"departmentIdType,omitempty"`
}
