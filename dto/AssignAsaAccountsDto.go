package dto


type AssignAsaAccountsDto struct{
    AppId  string `json:"appId"`
    AccountId  string `json:"accountId"`
    Targets  []AssignAsaAccountItem `json:"targets"`
}

