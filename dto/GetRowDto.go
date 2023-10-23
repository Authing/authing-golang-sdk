package dto


type GetRowDto struct{
    ModelId string `json:"modelId,omitempty"`
    RowId string `json:"rowId,omitempty"`
    ShowFieldId string `json:"showFieldId,omitempty"`
}

