package dto


type ListDto struct{
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
}

