package dto


type GetAdminAuditLogsDto struct{
    RequestId  string `json:"requestId,omitempty"`
    ClientIp  string `json:"clientIp,omitempty"`
    OperationType  string `json:"operationType,omitempty"`
    ResourceType  string `json:"resourceType,omitempty"`
    UserId  string `json:"userId,omitempty"`
    Success  bool `json:"success,omitempty"`
    Start  int `json:"start,omitempty"`
    End  int `json:"end,omitempty"`
    Pagination  ListWebhooksDto `json:"pagination,omitempty"`
}

