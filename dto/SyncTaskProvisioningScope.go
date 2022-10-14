package dto


type SyncTaskProvisioningScope struct{
    All  bool `json:"all"`
    IncludeNewUsers  bool `json:"includeNewUsers"`
}

