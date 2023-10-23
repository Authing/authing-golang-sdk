package dto


type WechatMobileAuthByCodeIdentityInput struct{
    Code  string `json:"code"`
    AppId  string `json:"appId,omitempty"`
    ConnId  string `json:"connId,omitempty"`
    Options  SignInOptionsDto `json:"options,omitempty"`
}

