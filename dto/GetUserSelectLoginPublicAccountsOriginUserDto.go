package dto


type GetUserSelectLoginPublicAccountsOriginUserDto struct{
    UserId  string `json:"userId"`
    Avatar  string `json:"avatar,omitempty"`
    DisplayName  string `json:"displayName,omitempty"`
    Usertype  string  `json:"usertype,omitempty"`
}

