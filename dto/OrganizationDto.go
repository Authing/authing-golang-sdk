package dto

type OrganizationDto struct {
	OrganizationCode string                  `json:"organizationCode"`
	OrganizationName string                  `json:"organizationName"`
	Description      string                  `json:"description,omitempty"`
	CreatedAt        string                  `json:"createdAt,omitempty"`
	UpdatedAt        string                  `json:"updatedAt,omitempty"`
	DepartmentId     string                  `json:"departmentId"`
	OpenDepartmentId string                  `json:"openDepartmentId,omitempty"`
	HasChildren      bool                    `json:"hasChildren"`
	LeaderUserIds    []string                `json:"leaderUserIds,omitempty"`
	MembersCount     int                     `json:"membersCount"`
	IsVirtualNode    bool                    `json:"isVirtualNode"`
	I18n             OrganizationNameI18nDto `json:"i18n,omitempty"`
	CustomData       interface{}             `json:"customData,omitempty"`
	TenantId         string                  `json:"tenantId,omitempty"`
}
