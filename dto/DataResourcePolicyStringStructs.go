package dto


type DataResourcePolicyStringStructs struct{
    OperationType  string `json:"operationType"`
    ActionList  []ArrayOrStringActionDto `json:"actionList"`
}

