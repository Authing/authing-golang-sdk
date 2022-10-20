package dto


type SignInByAdPayloadDto struct{
    Password  string `json:"password"`
    SAMAccountName  string `json:"sAMAccountName"`
}

