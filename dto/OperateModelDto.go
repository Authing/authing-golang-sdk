package dto


type OperateModelDto struct{
    Id  string `json:"id"`
    UserPoolId  string `json:"userPoolId"`
    ModelId  string `json:"modelId"`
    OperateName  string `json:"operateName"`
    OperateKey  string `json:"operateKey"`
    Show  bool `json:"show"`
    IsDefault  bool `json:"isDefault"`
    CreatedAt  string `json:"createdAt"`
    UpdatedAt  string `json:"updatedAt"`
}

