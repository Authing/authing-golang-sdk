package dto


type UpdateDepartmentReqDto struct{
    DepartmentId  string `json:"departmentId"`
    OrganizationCode  string `json:"organizationCode"`
    LeaderUserIds  []string `json:"leaderUserIds,omitempty"`
    Description  string `json:"description,omitempty"`
    Code  string `json:"code,omitempty"`
    I18n  DepartmentI18nDto `json:"i18n,omitempty"`
    Name  string `json:"name,omitempty"`
    DepartmentIdType  string  `json:"departmentIdType,omitempty"`
    ParentDepartmentId  string `json:"parentDepartmentId,omitempty"`
    CustomData  interface{} `json:"customData,omitempty"`
}

