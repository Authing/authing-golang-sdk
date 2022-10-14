package dto


type AccessTokenResDto struct{
    AccessToken  string `json:"access_token"`
    ExpiresIn  int `json:"expires_in"`
}

