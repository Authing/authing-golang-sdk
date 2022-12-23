package dto


type CreateAsaAccountsBatchDto struct{
    List  []CreateAsaAccountsBatchItemDto `json:"list"`
    AppId  string `json:"appId"`
}

