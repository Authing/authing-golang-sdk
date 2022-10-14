package dto


type AuthenticateByADDto struct{
    Password  string `json:"password"`
    SAMAccountName  string `json:"sAMAccountName"`
}

