package dto


type UpdateSyncTaskDto struct{
    SyncTaskId  int `json:"syncTaskId"`
    SyncTaskName  string `json:"syncTaskName,omitempty"`
    SyncTaskType  string  `json:"syncTaskType,omitempty"`
    ClientConfig  SyncTaskClientConfig `json:"clientConfig,omitempty"`
    SyncTaskFlow  string  `json:"syncTaskFlow,omitempty"`
    SyncTaskTrigger  string  `json:"syncTaskTrigger,omitempty"`
    OrganizationCode  string `json:"organizationCode,omitempty"`
    ProvisioningScope  SyncTaskProvisioningScope `json:"provisioningScope,omitempty"`
    FieldMapping  []SyncTaskFieldMapping `json:"fieldMapping,omitempty"`
    TimedScheduler  SyncTaskTimedScheduler `json:"timedScheduler,omitempty"`
}

