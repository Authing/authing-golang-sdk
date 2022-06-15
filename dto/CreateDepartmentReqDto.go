package dto

type CreateDepartmentReqDto struct {
	OrganizationCode   string  `json:"organizationCode"`
	ParentDepartmentId string  `json:"parentDepartmentId"`
	Name               string  `json:"name"`
	OpenDepartmentId   string  `json:"openDepartmentId,omitempty"`
	Description        string  `json:"description,omitempty"`
	Code               string  `json:"code,omitempty"`
	I18n               I18nDto `json:"i18n,omitempty"`
	DepartmentIdType   string  `json:"departmentIdType,omitempty"`
}
