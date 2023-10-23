package dto


type SigninByMobileDto struct{
    ExtIdpConnidentifier  string `json:"extIdpConnidentifier"`
    Connection  string  `json:"connection"`
    WechatPayload  SignInByWechatPayloadDto `json:"wechatPayload,omitempty"`
    ApplePayload  SignInByApplePayloadDto `json:"applePayload,omitempty"`
    AlipayPayload  SignInByAlipayPayloadDto `json:"alipayPayload,omitempty"`
    WechatworkPayload  SignInByWechatworkDto `json:"wechatworkPayload,omitempty"`
    WechatworkAgencyPayload  SignInByWechatworkAgencyPayloadDto `json:"wechatworkAgencyPayload,omitempty"`
    LarkPublicPayload  SignInByLarkPublicPayloadDto `json:"larkPublicPayload,omitempty"`
    LarkInternalPayload  SignInByLarkInternalPayloadDto `json:"larkInternalPayload,omitempty"`
    LarkBlockPayload  SignInByLarkBlockPayloadDto `json:"larkBlockPayload,omitempty"`
    YidunPayload  SignInByYidunPayloadDto `json:"yidunPayload,omitempty"`
    WechatMiniProgramCodePayload  SignInByWechatMiniProgramCodePayloadDto `json:"wechatMiniProgramCodePayload,omitempty"`
    WechatMiniProgramPhonePayload  SignInByWechatMiniProgramPhonePayloadDto `json:"wechatMiniProgramPhonePayload,omitempty"`
    WechatMiniProgramCodeAndPhonePayload  SignInByWechatMiniProgramCodeAndPhonePayloadDto `json:"wechatMiniProgramCodeAndPhonePayload,omitempty"`
    GooglePayload  SignInByGooglePayloadDto `json:"googlePayload,omitempty"`
    FacebookPayload  SignInByFacebookPayloadDto `json:"facebookPayload,omitempty"`
    QqPayload  SignInByQQPayloadDto `json:"qqPayload,omitempty"`
    WeiboPayload  SignInByWeiboPayloadDto `json:"weiboPayload,omitempty"`
    BaiduPayload  SignInByBaiduPayloadDto `json:"baiduPayload,omitempty"`
    LinkedInPayload  SignInByLinkedInPayloadDto `json:"linkedInPayload,omitempty"`
    DingTalkPayload  SignInByDingTalkPayloadDto `json:"dingTalkPayload,omitempty"`
    GithubPayload  SignInByGithubPayloadDto `json:"githubPayload,omitempty"`
    GiteePayload  SignInByGiteePayloadDto `json:"giteePayload,omitempty"`
    GitlabPayload  SignInByGitlabPayloadDto `json:"gitlabPayload,omitempty"`
    DouyinPayload  SignInByDouyinPayloadDto `json:"douyinPayload,omitempty"`
    KuaishouPayload  SignInByKuaishouPayloadDto `json:"kuaishouPayload,omitempty"`
    XiaomiPayload  SignInByXiaomiPayloadDto `json:"xiaomiPayload,omitempty"`
    LinePayload  SignInByLinePayloadDto `json:"linePayload,omitempty"`
    SlackPayload  SignInBySlackPayloadDto `json:"slackPayload,omitempty"`
    OppoPayload  SignInByOPPOPayloadDto `json:"oppoPayload,omitempty"`
    HuaweiPayload  SignInByHuaweiPayloadDto `json:"huaweiPayload,omitempty"`
    AmazonPayload  SignInByAmazonPayloadDto `json:"amazonPayload,omitempty"`
    Options  SignInByMobileOptionsDto `json:"options,omitempty"`
    ClientId  string `json:"client_id,omitempty"`
    ClientSecret  string `json:"client_secret,omitempty"`
}

