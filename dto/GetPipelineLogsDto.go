package dto


type GetPipelineLogsDto struct{
    FuncId string `json:"funcId,omitempty"`
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
}

