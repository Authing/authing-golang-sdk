package dto


type CreateOrganizationReqDto struct{
    OrganizationName  string `json:"organizationName"`
    OrganizationCode  string `json:"organizationCode"`
    OpenDepartmentId  string `json:"openDepartmentId,omitempty"`
}

