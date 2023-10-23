package dto


type CreateFunctionModelDto struct{
    ParentKey  string `json:"parentKey"`
    Enable  bool `json:"enable"`
    Type  string  `json:"type"`
    Description  string `json:"description"`
    Name  string `json:"name"`
    DataType  string  `json:"dataType,omitempty"`
}

