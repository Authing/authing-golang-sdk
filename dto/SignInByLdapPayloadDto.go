package dto


type SignInByLdapPayloadDto struct{
    Password  string `json:"password"`
    SAMAccountName  string `json:"sAMAccountName"`
}

