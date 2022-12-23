package dto


type CostGetOrderPayDetail struct{
    OrderNo  string `json:"orderNo"`
    ChannelOrderNo  string `json:"channelOrderNo"`
    PaidAmount  string `json:"paidAmount"`
    PaidTime  string `json:"paidTime"`
    PaidAccountNo  string `json:"paidAccountNo"`
    PayStatus  string `json:"payStatus"`
    CreateTime  string `json:"createTime"`
    PayType  string `json:"payType"`
}

