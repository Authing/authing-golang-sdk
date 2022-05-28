package dto


type GetUserLoginHistoryDto struct{
    UserId string `json:"user_id,omitempty"`
    AppId string `json:"app_id,omitempty"`
    ClientIp string `json:"client_ip,omitempty"`
    Start int `json:"start,omitempty"`
    End int `json:"end,omitempty"`
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
}

