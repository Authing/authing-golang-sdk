package dto


type AccessTokenDto struct{
    AccessToken  string `json:"access_token"`
    ExpiresIn  int `json:"expires_in"`
}

