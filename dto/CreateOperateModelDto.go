package dto


type CreateOperateModelDto struct{
    Show  bool `json:"show"`
    Icon  string `json:"icon"`
    Config  interface{} `json:"config"`
    OperateName  string `json:"operateName"`
    OperateKey  string `json:"operateKey"`
    ModelId  string `json:"modelId"`
}

