package dto

type UpdateOrganizationReqDto struct {
	OrganizationCode    string                  `json:"organizationCode"`
	Description         string                  `json:"description,omitempty"`
	OpenDepartmentId    string                  `json:"openDepartmentId,omitempty"`
	LeaderUserIds       []string                `json:"leaderUserIds,omitempty"`
	I18n                OrganizationNameI18nDto `json:"i18n,omitempty"`
	OrganizationNewCode string                  `json:"organizationNewCode,omitempty"`
	OrganizationName    string                  `json:"organizationName,omitempty"`
}
