package dto


type LoginTokenResponseDataDto struct{
    AccessToken  string `json:"access_token,omitempty"`
    IdToken  string `json:"id_token,omitempty"`
    RefreshToken  string `json:"refresh_token,omitempty"`
    TokenType  string `json:"token_type"`
    ExpireIn  int `json:"expire_in"`
}

