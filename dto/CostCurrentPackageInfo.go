package dto


type CostCurrentPackageInfo struct{
    Code  string `json:"code"`
    EndTime  string `json:"endTime"`
    OverdueDays  string `json:"overdueDays"`
    GoodsPackage  GoodsPackageDto `json:"goodsPackage"`
}

