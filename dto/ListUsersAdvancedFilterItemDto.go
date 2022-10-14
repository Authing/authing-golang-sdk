package dto


type ListUsersAdvancedFilterItemDto struct{
    Field  string `json:"field"`
    Operator  string  `json:"operator"`
    Value  interface{} `json:"value,omitempty"`
}

