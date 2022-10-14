package dto


type PaginationDto struct{
    Page  int `json:"page,omitempty"`
    Limit  int `json:"limit,omitempty"`
}

