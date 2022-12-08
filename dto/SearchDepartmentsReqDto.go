package dto


type SearchDepartmentsReqDto struct{
    OrganizationCode  string `json:"organizationCode"`
    Keywords  string `json:"keywords"`
    WithCustomData  bool `json:"withCustomData,omitempty"`
    Limit  int `json:"limit,omitempty"`
}

