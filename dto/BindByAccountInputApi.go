package dto


type BindByAccountInputApi struct{
    Key  string `json:"key"`
    Action  string `json:"action"`
    Password  string `json:"password"`
    Account  string `json:"account"`
}

