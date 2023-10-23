package dto


type IpListRespDto struct{
    Id  string `json:"id"`
    Ip  string `json:"ip"`
    IpType  string `json:"ipType"`
    AddType  string `json:"addType"`
    RemoveType  string `json:"removeType"`
    LimitList  []string `json:"limitList"`
}

