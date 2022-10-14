package dto


type GetUserDepartmentsDto struct{
    UserId string `json:"userId,omitempty"`
    UserIdType string `json:"userIdType,omitempty"`
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
    WithCustomData bool `json:"withCustomData,omitempty"`
    SortBy string `json:"sortBy,omitempty"`
    OrderBy string `json:"orderBy,omitempty"`
}

