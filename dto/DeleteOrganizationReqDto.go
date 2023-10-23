package dto


type DeleteOrganizationReqDto struct{
    OrganizationCode  string `json:"organizationCode"`
    TenantId  string `json:"tenantId,omitempty"`
}

