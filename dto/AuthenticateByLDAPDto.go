package dto


type AuthenticateByLDAPDto struct{
    Password  string `json:"password"`
    SAMAccountName  string `json:"sAMAccountName"`
}

