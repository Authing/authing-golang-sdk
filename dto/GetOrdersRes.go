package dto


type GetOrdersRes struct{
    TotalCount  string `json:"totalCount"`
    List  []OrderItem `json:"list"`
}

