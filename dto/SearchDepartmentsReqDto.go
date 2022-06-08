package dto

type SearchDepartmentsReqDto struct {
	Keywords         string `json:"keywords"`
	OrganizationCode string `json:"organizationCode"`
}
