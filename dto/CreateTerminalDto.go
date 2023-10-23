package dto


type CreateTerminalDto struct{
    DeviceUniqueId  string `json:"deviceUniqueId"`
    Type  string  `json:"type"`
    CustomData  interface{} `json:"customData"`
    Name  string `json:"name,omitempty"`
    Version  string `json:"version,omitempty"`
    Hks  string `json:"hks,omitempty"`
    Fde  string `json:"fde,omitempty"`
    Hor  bool `json:"hor,omitempty"`
    Sn  string `json:"sn,omitempty"`
    Producer  string `json:"producer,omitempty"`
    Mod  string `json:"mod,omitempty"`
    Os  string `json:"os,omitempty"`
    Imei  string `json:"imei,omitempty"`
    Meid  string `json:"meid,omitempty"`
    Description  string `json:"description,omitempty"`
    Language  string `json:"language,omitempty"`
    Cookie  bool `json:"cookie,omitempty"`
    UserAgent  string `json:"userAgent,omitempty"`
}

