package dto


type ListUsersRequestDto struct{
    Keywords  string `json:"keywords,omitempty"`
    AdvancedFilter  []ListUsersAdvancedFilterItemDto `json:"advancedFilter,omitempty"`
    SearchQuery  interface{} `json:"searchQuery,omitempty"`
    Options  ListUsersOptionsDto `json:"options,omitempty"`
}

