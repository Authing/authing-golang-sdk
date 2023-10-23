package dto


type LdapLogDto struct{
    TotalCount  int `json:"totalCount"`
    List  interface{} `json:"list"`
}

