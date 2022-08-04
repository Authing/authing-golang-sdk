package dto

type SearchDepartmentsReqDto struct {
	Keywords         string `json:"keywords"`
	OrganizationCode string `json:"organizationCode"`
	WithCustomData   bool   `json:"withCustomData,omitempty"`
}
