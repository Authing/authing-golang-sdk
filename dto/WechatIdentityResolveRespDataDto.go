package dto


type WechatIdentityResolveRespDataDto struct{
    Methods  []string `json:"methods,omitempty"`
    Accounts  []string `json:"accounts,omitempty"`
    Key  string `json:"key"`
}

