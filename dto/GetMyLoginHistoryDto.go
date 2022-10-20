package dto


type GetMyLoginHistoryDto struct{
    AppId string `json:"appId,omitempty"`
    ClientIp string `json:"clientIp,omitempty"`
    Success bool `json:"success,omitempty"`
    Start int `json:"start,omitempty"`
    End int `json:"end,omitempty"`
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
}

