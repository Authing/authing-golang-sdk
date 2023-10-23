package dto


type ListTenantsDto struct{
    Keywords string `json:"keywords,omitempty"`
    WithMembersCount bool `json:"withMembersCount,omitempty"`
    WithAppDetail bool `json:"withAppDetail,omitempty"`
    WithCreatorDetail bool `json:"withCreatorDetail,omitempty"`
    WithSourceAppDetail bool `json:"withSourceAppDetail,omitempty"`
    Page string `json:"page,omitempty"`
    Limit string `json:"limit,omitempty"`
    Source interface{} `json:"source,omitempty"`
}

