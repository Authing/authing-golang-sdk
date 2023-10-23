package dto


type ListTenantCooperatorsDto struct{
    Keywords string `json:"keywords,omitempty"`
    External bool `json:"external,omitempty"`
    Page string `json:"page,omitempty"`
    Limit string `json:"limit,omitempty"`
}

