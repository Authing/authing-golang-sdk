package dto


type SigninByMobileDto struct{
    ExtIdpConnidentifier  string `json:"extIdpConnidentifier"`
    Connection  string  `json:"connection"`
    WechatPayload  AuthenticateByWechatDto `json:"wechatPayload,omitempty"`
    AlipayPayload  AuthenticateByAlipayDto `json:"alipayPayload,omitempty"`
    WechatworkPayload  AuthenticateByWechatworkDto `json:"wechatworkPayload,omitempty"`
    WechatworkAgencyPayload  AuthenticateByWechatworkAgencyDto `json:"wechatworkAgencyPayload,omitempty"`
    LarkPublicPayload  AuthenticateByLarkPublicDto `json:"larkPublicPayload,omitempty"`
    LarkInternalPayload  AuthenticateByLarkInternalDto `json:"larkInternalPayload,omitempty"`
    YidunPayload  AuthenticateByYidunDto `json:"yidunPayload,omitempty"`
    WechatMiniProgramCodePayload  AuthenticateByWechatMiniProgramCodeDto `json:"wechatMiniProgramCodePayload,omitempty"`
    WechatMiniProgramPhonePayload  AuthenticateByWechatMiniProgramPhoneDto `json:"wechatMiniProgramPhonePayload,omitempty"`
    GooglePayload  AuthenticateByGoogleDto `json:"googlePayload,omitempty"`
    Options  SignInByMobileOptionsDto `json:"options,omitempty"`
    ClientId  string `json:"client_id,omitempty"`
    ClientSecret  string `json:"client_secret,omitempty"`
}

