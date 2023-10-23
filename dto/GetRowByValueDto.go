package dto


type GetRowByValueDto struct{
    ModelId string `json:"modelId,omitempty"`
    Key string `json:"key,omitempty"`
    Value string `json:"value,omitempty"`
    ShowFieldId string `json:"showFieldId,omitempty"`
}

