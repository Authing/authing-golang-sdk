package dto


type CreateUEBADto struct{
    Data  interface{} `json:"data"`
    ModelId  string `json:"modelId,omitempty"`
}

