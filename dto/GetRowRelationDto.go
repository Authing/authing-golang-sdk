package dto


type GetRowRelationDto struct{
    ModelId string `json:"modelId,omitempty"`
    FieldId string `json:"fieldId,omitempty"`
    RowId string `json:"rowId,omitempty"`
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
}

