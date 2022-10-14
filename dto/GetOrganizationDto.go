package dto


type GetOrganizationDto struct{
    OrganizationCode string `json:"organizationCode,omitempty"`
    WithCustomData bool `json:"withCustomData,omitempty"`
}

