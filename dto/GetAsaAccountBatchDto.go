package dto


type GetAsaAccountBatchDto struct{
    AccountIds  []string `json:"accountIds"`
    AppId  string `json:"appId"`
}

