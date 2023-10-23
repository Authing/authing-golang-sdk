package dto


type RiskListPolicyRespDto struct{
    Id  string `json:"id"`
    OptObject  string  `json:"optObject"`
    IpRange  string  `json:"ipRange"`
    UserRange  string  `json:"userRange"`
    IpCond  string  `json:"ipCond"`
    UserCond  string  `json:"userCond"`
    TimeRange  int `json:"timeRange"`
    CountThr  int `json:"countThr"`
    EventStateType  string  `json:"eventStateType"`
    RemoveType  string  `json:"removeType"`
    Action  string  `json:"action"`
    LimitList  string  `json:"limitList"`
    CreatedAt  string `json:"createdAt"`
}

