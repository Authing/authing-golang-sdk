package dto


type GetOrganizationDto struct{
    OrganizationCode string `json:"organizationCode,omitempty"`
    WithCustomData bool `json:"withCustomData,omitempty"`
    WithPost bool `json:"withPost,omitempty"`
    TenantId string `json:"tenantId,omitempty"`
}

