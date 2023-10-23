package dto


type GetUserSelectLoginPublicAccountsDataDto struct{
    PublicAccounts  []ListPublicAccountDataDto `json:"publicAccounts"`
    OriginUser  GetUserSelectLoginPublicAccountsOriginUserDto `json:"originUser,omitempty"`
}

