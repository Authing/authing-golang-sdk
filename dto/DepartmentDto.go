package dto

type DepartmentDto struct {
	OrganizationCode   string      `json:"organizationCode"`
	DepartmentId       string      `json:"departmentId"`
	CreatedAt          string      `json:"createdAt"`
	OpenDepartmentId   string      `json:"openDepartmentId,omitempty"`
	Name               string      `json:"name"`
	LeaderUserIds      []string    `json:"leaderUserIds,omitempty"`
	Description        string      `json:"description,omitempty"`
	ParentDepartmentId string      `json:"parentDepartmentId"`
	Code               string      `json:"code,omitempty"`
	MembersCount       int         `json:"membersCount"`
	HasChildren        bool        `json:"hasChildren"`
	IsVirtualNode      bool        `json:"isVirtualNode,omitempty"`
	I18n               I18nDto     `json:"i18n,omitempty"`
	CustomData         interface{} `json:"customData,omitempty"`
}
