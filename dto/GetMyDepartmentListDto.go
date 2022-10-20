package dto


type GetMyDepartmentListDto struct{
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
    WithCustomData bool `json:"withCustomData,omitempty"`
    SortBy string `json:"sortBy,omitempty"`
    OrderBy string `json:"orderBy,omitempty"`
}

