package dto


type CreateRowDto struct{
    Data  interface{} `json:"data"`
    ModelId  string `json:"modelId"`
    RowId  string `json:"rowId,omitempty"`
}

