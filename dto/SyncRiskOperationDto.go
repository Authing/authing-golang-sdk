package dto


type SyncRiskOperationDto struct{
    SyncRiskOperationId  int `json:"syncRiskOperationId"`
    SyncTaskId  int `json:"syncTaskId"`
    CreatedAt  string `json:"createdAt"`
    UpdatedAt  string `json:"updatedAt"`
    Status  string  `json:"status"`
    Level  int `json:"level"`
    ObjectType  string  `json:"objectType"`
    ObjectName  string `json:"objectName"`
    ObjectId  string `json:"objectId"`
    ErrMsg  string `json:"errMsg,omitempty"`
}

