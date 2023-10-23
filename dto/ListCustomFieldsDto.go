package dto


type ListCustomFieldsDto struct{
    TargetType string `json:"targetType,omitempty"`
    DataType string `json:"dataType,omitempty"`
    TenantId string `json:"tenantId,omitempty"`
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
    UserVisible bool `json:"userVisible,omitempty"`
    AdminVisible bool `json:"adminVisible,omitempty"`
    AccessControl bool `json:"accessControl,omitempty"`
    Keyword string `json:"keyword,omitempty"`
    Lang string `json:"lang,omitempty"`
}

