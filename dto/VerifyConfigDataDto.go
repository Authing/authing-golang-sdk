package dto


type VerifyConfigDataDto struct{
    Id  string `json:"id"`
    AppId  string `json:"appId"`
    IsWinLogin  bool `json:"isWinLogin,omitempty"`
    IsWinChangePwd  bool `json:"isWinChangePwd,omitempty"`
    AppSecret  string `json:"appSecret"`
    PublicKey  string `json:"publicKey"`
    AuthAddress  string `json:"authAddress"`
    Logo  string `json:"logo"`
    Name  string `json:"name"`
}

