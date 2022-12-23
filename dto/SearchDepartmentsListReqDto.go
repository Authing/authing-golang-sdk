package dto


type SearchDepartmentsListReqDto struct{
    OrganizationCode  string `json:"organizationCode"`
    WithCustomData  bool `json:"withCustomData,omitempty"`
    Page  int `json:"page,omitempty"`
    Limit  int `json:"limit,omitempty"`
    AdvancedFilter  []SearchDepartmentsFilterItemDto `json:"advancedFilter,omitempty"`
    SortBy  string  `json:"sortBy,omitempty"`
    OrderBy  string  `json:"orderBy,omitempty"`
}

