package dto


type UpdateRowDto struct{
    Data  interface{} `json:"data"`
    RowId  string `json:"rowId"`
    ModelId  string `json:"modelId"`
    ShowFieldId  bool `json:"showFieldId,omitempty"`
}

