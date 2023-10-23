package dto


type IpListCreateDto struct{
    ExpireAt  string `json:"expireAt"`
    LimitList  []string `json:"limitList"`
    RemoveType  string `json:"removeType"`
    AddType  string `json:"addType"`
    IpType  string `json:"ipType"`
    Ips  string `json:"ips"`
}

