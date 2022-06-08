package dto

type GetUserLoginHistoryDto struct {
	UserId   string `json:"userId,omitempty"`
	AppId    string `json:"appId,omitempty"`
	ClientIp string `json:"clientIp,omitempty"`
	Start    int    `json:"start,omitempty"`
	End      int    `json:"end,omitempty"`
	Page     int    `json:"page,omitempty"`
	Limit    int    `json:"limit,omitempty"`
}
