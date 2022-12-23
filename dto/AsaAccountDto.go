package dto


type AsaAccountDto struct{
    AppId  string `json:"appId"`
    AccountId  string `json:"accountId"`
    AccountInfo  interface{} `json:"accountInfo"`
}

