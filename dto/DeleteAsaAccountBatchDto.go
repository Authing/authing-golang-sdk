package dto


type DeleteAsaAccountBatchDto struct{
    AccountIds  []string `json:"accountIds"`
    AppId  string `json:"appId"`
}

