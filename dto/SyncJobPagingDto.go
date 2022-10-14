package dto


type SyncJobPagingDto struct{
    TotalCount  int `json:"totalCount"`
    List  []SyncJobDto `json:"list"`
}

