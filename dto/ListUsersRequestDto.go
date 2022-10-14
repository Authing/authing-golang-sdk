package dto


type ListUsersRequestDto struct{
    Query  string `json:"query,omitempty"`
    AdvancedFilter  []ListUsersAdvancedFilterItemDto `json:"advancedFilter,omitempty"`
    Options  ListUsersOptionsDto `json:"options,omitempty"`
}

