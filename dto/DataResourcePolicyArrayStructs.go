package dto


type DataResourcePolicyArrayStructs struct{
    OperationType  string `json:"operationType"`
    ActionList  []ArrayOrStringActionDto `json:"actionList"`
}

