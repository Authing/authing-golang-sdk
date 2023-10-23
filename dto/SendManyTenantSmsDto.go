package dto


type SendManyTenantSmsDto struct{
    AdminName  string `json:"adminName"`
    ImportId  int `json:"importId"`
    Users  []SendTenantSmsDto `json:"users"`
}

