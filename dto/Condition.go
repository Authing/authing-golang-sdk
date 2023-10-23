package dto


type Condition struct{
    Key  string `json:"key"`
    Value  interface{} `json:"value"`
    Operator  string `json:"operator"`
}

