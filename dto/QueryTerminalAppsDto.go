package dto


type QueryTerminalAppsDto struct{
    DeviceIds  []string `json:"deviceIds"`
    UserId  string `json:"userId,omitempty"`
}

