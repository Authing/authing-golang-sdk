package dto


type GetRowBatchDto struct{
    RowIds  []string `json:"rowIds"`
    ModelId  string `json:"modelId"`
}

