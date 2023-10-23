package dto


type ListPublicAccountsOptionsDto struct{
    Pagination  PaginationDto `json:"pagination,omitempty"`
    Sort  []SortingDto `json:"sort,omitempty"`
    FuzzySearchOn  []string `json:"fuzzySearchOn,omitempty"`
    WithCustomData  bool `json:"withCustomData,omitempty"`
    WithDepartmentIds  bool `json:"withDepartmentIds,omitempty"`
}

