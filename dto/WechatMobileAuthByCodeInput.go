package dto


type WechatMobileAuthByCodeInput struct{
    Code  string `json:"code"`
    AppId  string `json:"appId,omitempty"`
    ConnId  string `json:"connId,omitempty"`
}

