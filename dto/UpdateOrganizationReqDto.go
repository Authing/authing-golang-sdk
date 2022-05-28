package dto


type UpdateOrganizationReqDto struct{
    OrganizationCode  string `json:"organizationCode"`
    OpenDepartmentId  string `json:"openDepartmentId,omitempty"`
    OrganizationNewCode  string `json:"organizationNewCode,omitempty"`
    OrganizationName  string `json:"organizationName,omitempty"`
}

