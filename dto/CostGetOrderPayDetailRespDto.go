package dto


type CostGetOrderPayDetailRespDto struct{
    StatusCode  int `json:"statusCode"`
    Message  string `json:"message"`
    ApiCode  int `json:"apiCode,omitempty"`
    RequestId  string `json:"requestId,omitempty"`
    OrderNo  string `json:"orderNo"`
    ChannelOrderNo  string `json:"channelOrderNo"`
    PaidAmount  string `json:"paidAmount"`
    PaidTime  string `json:"paidTime"`
    PaidAccountNo  string `json:"paidAccountNo"`
    PayStatus  string `json:"payStatus"`
    CreateTime  string `json:"createTime"`
    PayType  string `json:"payType"`
}

