package dto


type SigninByMobileDto struct{
    ExtIdpConnidentifier  string `json:"extIdpConnidentifier"`
    Connection  string  `json:"connection"`
    WechatPayload  SignInByWechatPayloadDto `json:"wechatPayload,omitempty"`
    AlipayPayload  SignInByAlipayPayloadDto `json:"alipayPayload,omitempty"`
    WechatworkPayload  AuthenticateByWechatworkDto `json:"wechatworkPayload,omitempty"`
    WechatworkAgencyPayload  SignInByWechatworkAgencyPayloadDto `json:"wechatworkAgencyPayload,omitempty"`
    LarkPublicPayload  SignInByLarkPublicPayloadDto `json:"larkPublicPayload,omitempty"`
    LarkInternalPayload  SignInByLarkInternalPayloadDto `json:"larkInternalPayload,omitempty"`
    YidunPayload  SignInByYidunPayloadDto `json:"yidunPayload,omitempty"`
    WechatMiniProgramCodePayload  SignInByWechatMiniProgramCodePayloadDto `json:"wechatMiniProgramCodePayload,omitempty"`
    WechatMiniProgramPhonePayload  SignInByWechatMiniProgramPhonePayloadDto `json:"wechatMiniProgramPhonePayload,omitempty"`
    GooglePayload  SignInByGooglePayloadDto `json:"googlePayload,omitempty"`
    Options  SignInByMobileOptionsDto `json:"options,omitempty"`
    ClientId  string `json:"client_id,omitempty"`
    ClientSecret  string `json:"client_secret,omitempty"`
}

