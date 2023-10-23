package dto


type GetCustomFieldsDto struct{
    TargetType string `json:"targetType,omitempty"`
    TenantId string `json:"tenantId,omitempty"`
}

