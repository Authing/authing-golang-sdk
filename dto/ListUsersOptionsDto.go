package dto


type ListUsersOptionsDto struct{
    Pagination  PaginationDto `json:"pagination,omitempty"`
    Sort  []SortingDto `json:"sort,omitempty"`
    FuzzySearchOn  []string `json:"fuzzySearchOn,omitempty"`
    WithCustomData  bool `json:"withCustomData,omitempty"`
    WithPost  bool `json:"withPost,omitempty"`
    WithIdentities  bool `json:"withIdentities,omitempty"`
    WithDepartmentIds  bool `json:"withDepartmentIds,omitempty"`
    FlatCustomData  bool `json:"flatCustomData,omitempty"`
}

