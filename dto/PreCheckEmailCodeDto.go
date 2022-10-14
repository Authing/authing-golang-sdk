package dto


type PreCheckEmailCodeDto struct{
    Email  string `json:"email"`
    PassCode  string `json:"passCode"`
    Channel  string  `json:"channel"`
}

