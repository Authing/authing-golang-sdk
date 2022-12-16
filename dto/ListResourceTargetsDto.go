package dto


type ListResourceTargetsDto struct{
    Resource  string `json:"resource"`
    ActionAuthList  []ActionAuth `json:"actionAuthList"`
}

