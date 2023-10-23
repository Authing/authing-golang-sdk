package dto


type UpdateFunctionModelDto struct{
    Config  interface{} `json:"config"`
    FieldOrder  string `json:"fieldOrder"`
    Type  string  `json:"type"`
    ParentKey  string `json:"parentKey"`
    Enable  bool `json:"enable"`
    Description  string `json:"description"`
    Name  string `json:"name"`
    Id  string `json:"id"`
}

