package dto


type ListSyncRiskOperationsDto struct{
    SyncTaskId int `json:"syncTaskId,omitempty"`
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
    Status string `json:"status,omitempty"`
    ObjectType string `json:"objectType,omitempty"`
}

