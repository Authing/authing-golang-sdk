package dto


type WechatIdentityResolveRespDto struct{
    UniqueId  string `json:"uniqueId,omitempty"`
    Code  int `json:"code,omitempty"`
    StatusCode  int `json:"statusCode"`
    ApiCode  int `json:"apiCode,omitempty"`
    Message  string `json:"message"`
    Data  WechatIdentityResolveRespDataDto `json:"data"`
}

