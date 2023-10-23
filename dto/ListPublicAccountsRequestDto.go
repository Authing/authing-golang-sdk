package dto


type ListPublicAccountsRequestDto struct{
    Keywords  string `json:"keywords,omitempty"`
    AdvancedFilter  []ListPublicAccountsAdvancedFilterItemDto `json:"advancedFilter,omitempty"`
    SearchQuery  interface{} `json:"searchQuery,omitempty"`
    Options  ListPublicAccountsOptionsDto `json:"options,omitempty"`
}

