package dto


type GetWechatAccessTokenDataDto struct{
    AccessToken  string `json:"accessToken"`
    ExpiresAt  int `json:"expiresAt"`
}

