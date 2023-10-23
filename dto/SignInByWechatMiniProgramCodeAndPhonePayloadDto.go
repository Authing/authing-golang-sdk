package dto


type SignInByWechatMiniProgramCodeAndPhonePayloadDto struct{
    WxLoginInfo  SignInByWechatMiniProgramCodePayloadDto `json:"wxLoginInfo"`
    WxPhoneInfo  SignInByWechatMiniProgramPhoneInfoPayloadDto `json:"wxPhoneInfo"`
}

