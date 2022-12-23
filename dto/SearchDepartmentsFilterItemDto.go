package dto


type SearchDepartmentsFilterItemDto struct{
    Field  string `json:"field"`
    Operator  string  `json:"operator"`
    Value  interface{} `json:"value,omitempty"`
}

