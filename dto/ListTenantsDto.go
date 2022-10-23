package dto


type ListTenantsDto struct{
    Keywords string `json:"keywords,omitempty"`
    WithMembersCount bool `json:"withMembersCount,omitempty"`
    WithAppDetail bool `json:"withAppDetail,omitempty"`
    Page string `json:"page,omitempty"`
    Limit string `json:"limit,omitempty"`
}

