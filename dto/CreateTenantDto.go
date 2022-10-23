package dto


type CreateTenantDto struct{
    RejectHint  string `json:"rejectHint"`
    AppIds  []string `json:"appIds"`
    Name  string `json:"name"`
    Logo  []string `json:"logo,omitempty"`
    Description  string `json:"description,omitempty"`
}

