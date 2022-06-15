package dto

type DepartmentDto struct {
	DepartmentId       string   `json:"departmentId"`
	OpenDepartmentId   string   `json:"openDepartmentId,omitempty"`
	Name               string   `json:"name"`
	LeaderUserIds      []string `json:"leaderUserIds,omitempty"`
	Description        string   `json:"description,omitempty"`
	ParentDepartmentId string   `json:"parentDepartmentId"`
	Code               string   `json:"code,omitempty"`
	MembersCount       int      `json:"membersCount"`
	HasChildren        bool     `json:"hasChildren"`
	I18n               I18nDto  `json:"i18n,omitempty"`
}
