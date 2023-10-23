package dto


type CreateRelationValueDto struct{
    ValueList  []string `json:"valueList"`
    RowId  string `json:"rowId"`
    FieldId  string `json:"fieldId"`
    ModelId  string `json:"modelId"`
}

