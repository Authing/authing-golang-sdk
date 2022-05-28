package dto


type OrganizationDto struct{
    OrganizationCode  string `json:"organizationCode"`
    OrganizationName  string `json:"organizationName"`
    DepartmentId  string `json:"departmentId"`
    OpenDepartmentId  string `json:"openDepartmentId,omitempty"`
    HasChildren  bool `json:"hasChildren"`
    LeaderUserId  string `json:"leaderUserId"`
    MembersCount  int `json:"membersCount"`
}

