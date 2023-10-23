package dto


type FunctionModelDto struct{
    Id  string `json:"id"`
    UserPoolId  string `json:"userPoolId"`
    Name  string `json:"name"`
    Description  string `json:"description"`
    DataType  string  `json:"dataType"`
    Enable  bool `json:"enable"`
    ParentKey  string `json:"parentKey"`
    CreatedAt  string `json:"createdAt"`
    UpdatedAt  string `json:"updatedAt"`
    Type  string  `json:"type"`
    FieldOrder  string `json:"fieldOrder"`
    Config  interface{} `json:"config"`
}

