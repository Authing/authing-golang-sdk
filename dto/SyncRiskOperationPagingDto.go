package dto


type SyncRiskOperationPagingDto struct{
    TotalCount  int `json:"totalCount"`
    List  []SyncRiskOperationDto `json:"list"`
}

