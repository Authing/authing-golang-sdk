package dto


type RiskListPolicyCreateDto struct{
    LimitList  string  `json:"limitList"`
    Action  string  `json:"action"`
    RemoveType  string  `json:"removeType"`
    EventStateType  string  `json:"eventStateType"`
    CountThr  int `json:"countThr"`
    TimeRange  int `json:"timeRange"`
    UserCond  string  `json:"userCond"`
    IpCond  string  `json:"ipCond"`
    UserRange  string  `json:"userRange"`
    IpRange  string  `json:"ipRange"`
    OptObject  string  `json:"optObject"`
}

