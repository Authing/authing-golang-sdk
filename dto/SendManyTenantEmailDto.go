package dto


type SendManyTenantEmailDto struct{
    AdminName  string `json:"adminName"`
    ImportId  int `json:"importId"`
    Users  []SendTenantEmailDto `json:"users"`
}

