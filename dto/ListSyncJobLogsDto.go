package dto


type ListSyncJobLogsDto struct{
    SyncJobId int `json:"syncJobId,omitempty"`
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
    Success bool `json:"success,omitempty"`
    Action string `json:"action,omitempty"`
    ObjectType string `json:"objectType,omitempty"`
}

