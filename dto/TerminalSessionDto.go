package dto


type TerminalSessionDto struct{
    Device  DeviceInfo `json:"device"`
    LastLoginTime  string `json:"lastLoginTime"`
    LastIp  string `json:"lastIp,omitempty"`
    Online  bool `json:"online"`
}

