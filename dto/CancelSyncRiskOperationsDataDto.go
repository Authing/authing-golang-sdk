package dto


type CancelSyncRiskOperationsDataDto struct{
    SuccessList  []int `json:"successList"`
    FaildList  []int `json:"faildList"`
}

