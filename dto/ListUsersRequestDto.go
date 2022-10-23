package dto


type ListUsersRequestDto struct{
    Keywords  string `json:"keywords,omitempty"`
    AdvancedFilter  []ListUsersAdvancedFilterItemDto `json:"advancedFilter,omitempty"`
    Options  ListUsersOptionsDto `json:"options,omitempty"`
}

