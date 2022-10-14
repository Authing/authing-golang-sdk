package dto


type SyncTaskPagingDto struct{
    TotalCount  int `json:"totalCount"`
    List  []SyncTaskDto `json:"list"`
}

