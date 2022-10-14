package dto


type GetUserActionLogsDto struct{
    RequestId  string `json:"requestId,omitempty"`
    ClientIp  string `json:"clientIp,omitempty"`
    EventType  string `json:"eventType,omitempty"`
    UserId  string `json:"userId,omitempty"`
    AppId  string `json:"appId,omitempty"`
    Start  int `json:"start,omitempty"`
    End  int `json:"end,omitempty"`
    Success  bool `json:"success,omitempty"`
    Pagination  ListWebhooksDto `json:"pagination,omitempty"`
}

