package dto


type BindByEmailCodeInputApi struct{
    Key  string `json:"key"`
    Action  string `json:"action"`
    Code  string `json:"code"`
    Email  string `json:"email"`
}

