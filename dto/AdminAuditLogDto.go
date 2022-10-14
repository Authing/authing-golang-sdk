package dto


type AdminAuditLogDto struct{
    AdminUserId  string `json:"adminUserId"`
    AdminUserAvatar  string `json:"adminUserAvatar"`
    AdminUserDisplayName  string `json:"adminUserDisplayName"`
    ClientIp  string `json:"clientIp,omitempty"`
    OperationType  string  `json:"operationType"`
    ResourceType  string  `json:"resourceType"`
    EventDetail  string `json:"eventDetail,omitempty"`
    OperationParam  string `json:"operationParam,omitempty"`
    OriginValue  string `json:"originValue,omitempty"`
    TargetValue  string `json:"targetValue,omitempty"`
    Success  bool `json:"success"`
    UserAgent  string `json:"userAgent"`
    ParsedUserAgent  ParsedUserAgent `json:"parsedUserAgent"`
    Geoip  GeoIp `json:"geoip"`
    Timestamp  string `json:"timestamp"`
    RequestId  string `json:"requestId"`
}

