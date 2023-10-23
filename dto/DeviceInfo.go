package dto


type DeviceInfo struct{
    DeviceId  string `json:"deviceId"`
    Name  string `json:"name,omitempty"`
    Version  string `json:"version,omitempty"`
    Type  string  `json:"type"`
    Mod  string `json:"mod,omitempty"`
    Os  string `json:"os,omitempty"`
    Status  string  `json:"status,omitempty"`
    UserAgent  string `json:"userAgent,omitempty"`
}

