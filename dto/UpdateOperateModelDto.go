package dto


type UpdateOperateModelDto struct{
    Icon  string `json:"icon"`
    Config  interface{} `json:"config"`
    OperateName  string `json:"operateName"`
    OperateKey  string `json:"operateKey"`
    Show  bool `json:"show"`
    ModelId  string `json:"modelId"`
    Id  string `json:"id"`
}

