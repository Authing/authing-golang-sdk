package dto


type CreateDepartmentReqDto struct{
    OrganizationCode  string `json:"organizationCode"`
    Name  string `json:"name"`
    ParentDepartmentId  string `json:"parentDepartmentId"`
    Metadata  interface{} `json:"metadata"`
    OpenDepartmentId  string `json:"openDepartmentId,omitempty"`
    Description  string `json:"description,omitempty"`
    Code  string `json:"code,omitempty"`
    IsVirtualNode  bool `json:"isVirtualNode,omitempty"`
    I18n  DepartmentI18nDto `json:"i18n,omitempty"`
    CustomData  interface{} `json:"customData,omitempty"`
    DepartmentIdType  string  `json:"departmentIdType,omitempty"`
    PostIdList  []string `json:"postIdList,omitempty"`
    TenantId  string `json:"tenantId,omitempty"`
}

