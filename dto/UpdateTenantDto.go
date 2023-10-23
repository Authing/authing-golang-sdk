package dto


type UpdateTenantDto struct{
    TenantId  string `json:"tenantId"`
    Name  string `json:"name,omitempty"`
    AppIds  []string `json:"appIds,omitempty"`
    Logo  []string `json:"logo,omitempty"`
    Description  string `json:"description,omitempty"`
    RejectHint  string `json:"rejectHint,omitempty"`
    SourceAppId  string `json:"sourceAppId,omitempty"`
}

