package dto


type ListAsaAccountsDto struct{
    AppId string `json:"appId,omitempty"`
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
}

