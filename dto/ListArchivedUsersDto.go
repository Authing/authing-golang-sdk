package dto


type ListArchivedUsersDto struct{
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
    StartAt int `json:"startAt,omitempty"`
}

