package dto


type ListEventsDto struct{
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
    App string `json:"app,omitempty"`
}

