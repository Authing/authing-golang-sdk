package dto


type ListSyncJobsDto struct{
    SyncTaskId int `json:"syncTaskId,omitempty"`
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
    SyncTrigger string `json:"syncTrigger,omitempty"`
}

