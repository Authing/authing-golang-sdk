package dto


type ListPostDto struct{
    Keywords string `json:"keywords,omitempty"`
    SkipCount bool `json:"skipCount,omitempty"`
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
    WithMetadata bool `json:"withMetadata,omitempty"`
    WithCustomData bool `json:"withCustomData,omitempty"`
    FlatCustomData bool `json:"flatCustomData,omitempty"`
}

