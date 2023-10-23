package dto


type ListMultipleTenantAdminsDto struct{
    Keywords string `json:"keywords,omitempty"`
    Page string `json:"page,omitempty"`
    Limit string `json:"limit,omitempty"`
}

