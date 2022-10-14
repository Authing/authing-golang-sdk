package dto


type UserDepartmentRespDto struct{
    OrganizationCode  string `json:"organizationCode"`
    DepartmentId  string `json:"departmentId"`
    CreatedAt  string `json:"createdAt"`
    Name  string `json:"name"`
    Description  string `json:"description"`
    OpenDepartmentId  string `json:"openDepartmentId,omitempty"`
    IsLeader  bool `json:"isLeader"`
    Code  string `json:"code"`
    IsMainDepartment  bool `json:"isMainDepartment"`
    JoinedAt  string `json:"joinedAt"`
    IsVirtualNode  bool `json:"isVirtualNode"`
    I18n  DepartmentI18nDto `json:"i18n,omitempty"`
    CustomData  interface{} `json:"customData,omitempty"`
}

