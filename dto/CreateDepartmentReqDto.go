package dto


type CreateDepartmentReqDto struct{
    ParentDepartmentId  string `json:"parentDepartmentId"`
    Name  string `json:"name"`
    OrganizationCode  string `json:"organizationCode"`
    OpenDepartmentId  string `json:"openDepartmentId,omitempty"`
    Description  string `json:"description,omitempty"`
    Code  string `json:"code,omitempty"`
    IsVirtualNode  bool `json:"isVirtualNode,omitempty"`
    I18n  DepartmentI18nDto `json:"i18n,omitempty"`
    CustomData  interface{} `json:"customData,omitempty"`
    DepartmentIdType  string  `json:"departmentIdType,omitempty"`
}

