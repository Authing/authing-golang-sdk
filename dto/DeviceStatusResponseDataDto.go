package dto


type DeviceStatusResponseDataDto struct{
    Status  string  `json:"status"`
    DiffTime  int `json:"diffTime,omitempty"`
}

