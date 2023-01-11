package dto

type CreateOrganizationReqDto struct {
	OrganizationName string                  `json:"organizationName"`
	OrganizationCode string                  `json:"organizationCode"`
	Description      string                  `json:"description,omitempty"`
	OpenDepartmentId string                  `json:"openDepartmentId,omitempty"`
	I18n             OrganizationNameI18nDto `json:"i18n,omitempty"`
	TenantId         string                  `json:"tenantId,omitempty"`
}
