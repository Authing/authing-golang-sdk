package dto


type RemoveRelationValueDto struct{
    Value  string `json:"value"`
    FieldIds  []string `json:"fieldIds"`
    RowId  string `json:"rowId"`
    ModelId  string `json:"modelId"`
}

