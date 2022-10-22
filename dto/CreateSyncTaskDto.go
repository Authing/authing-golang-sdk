package dto


type CreateSyncTaskDto struct{
    FieldMapping  []SyncTaskFieldMapping `json:"fieldMapping"`
    SyncTaskTrigger  string  `json:"syncTaskTrigger"`
    SyncTaskFlow  string  `json:"syncTaskFlow"`
    ClientConfig  SyncTaskClientConfig `json:"clientConfig"`
    SyncTaskType  string  `json:"syncTaskType"`
    SyncTaskName  string `json:"syncTaskName"`
    OrganizationCode  string `json:"organizationCode,omitempty"`
    ProvisioningScope  SyncTaskProvisioningScope `json:"provisioningScope,omitempty"`
    TimedScheduler  SyncTaskTimedScheduler `json:"timedScheduler,omitempty"`
}

