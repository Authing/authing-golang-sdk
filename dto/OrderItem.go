package dto


type OrderItem struct{
    OrderNo  string `json:"orderNo"`
    GoodsName  string `json:"goodsName"`
    GoodsNameEn  string `json:"goodsNameEn"`
    GoodsUnitPrice  string `json:"goodsUnitPrice"`
    Quantity  string `json:"quantity"`
    ActualAmount  string `json:"actualAmount"`
    Status  string `json:"status"`
    OrderType  string `json:"orderType"`
    CreateTime  string `json:"createTime"`
    Source  string `json:"source"`
}

