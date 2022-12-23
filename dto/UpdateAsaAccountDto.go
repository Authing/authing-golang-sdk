package dto


type UpdateAsaAccountDto struct{
    AccountInfo  interface{} `json:"accountInfo"`
    AccountId  string `json:"accountId"`
    AppId  string `json:"appId"`
}

