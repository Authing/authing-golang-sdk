package dto


type ListOperateDto struct{
    ModelId string `json:"modelId,omitempty"`
    Keywords string `json:"keywords,omitempty"`
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
}

