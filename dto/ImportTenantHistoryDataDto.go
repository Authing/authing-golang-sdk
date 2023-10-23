package dto


type ImportTenantHistoryDataDto struct{
    TotalCount  int `json:"totalCount"`
    List  []interface{} `json:"list"`
}

