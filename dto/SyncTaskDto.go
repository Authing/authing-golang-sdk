package dto


type SyncTaskDto struct{
    SyncTaskId  int `json:"syncTaskId"`
    CreatedAt  string `json:"createdAt"`
    UpdatedAt  string `json:"updatedAt"`
    SyncTaskName  string `json:"syncTaskName"`
    SyncTaskType  string  `json:"syncTaskType"`
    SyncFlow  string  `json:"syncFlow"`
    SyncTrigger  string  `json:"syncTrigger"`
    LastSyncMessage  bool `json:"lastSyncMessage,omitempty"`
    LastSyncRate  int `json:"lastSyncRate,omitempty"`
    LastSyncStatus  string  `json:"lastSyncStatus,omitempty"`
    LastSyncTime  string `json:"lastSyncTime,omitempty"`
    OrganizationCode  string `json:"organizationCode,omitempty"`
    ProvisioningScope  SyncTaskProvisioningScope `json:"provisioningScope,omitempty"`
    FieldMapping  []SyncTaskFieldMapping `json:"fieldMapping"`
    TimedScheduler  SyncTaskTimedScheduler `json:"timedScheduler,omitempty"`
}

