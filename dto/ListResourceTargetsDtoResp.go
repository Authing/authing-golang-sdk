package dto


type ListResourceTargetsDtoResp struct{
    Resource  string `json:"resource"`
    ActionAuthList  []ActionAuth `json:"actionAuthList"`
}

