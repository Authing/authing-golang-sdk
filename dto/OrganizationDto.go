package dto

type OrganizationDto struct {
	OrganizationCode string                  `json:"organizationCode"`
	OrganizationName string                  `json:"organizationName"`
	Description      string                  `json:"description,omitempty"`
	DepartmentId     string                  `json:"departmentId"`
	OpenDepartmentId string                  `json:"openDepartmentId,omitempty"`
	HasChildren      bool                    `json:"hasChildren"`
	LeaderUserIds    []string                `json:"leaderUserIds,omitempty"`
	MembersCount     int                     `json:"membersCount"`
	I18n             OrganizationNameI18nDto `json:"i18n,omitempty"`
}
