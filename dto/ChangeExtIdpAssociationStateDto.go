package dto


type ChangeExtIdpAssociationStateDto struct{
    Association  bool `json:"association"`
    Id  string `json:"id"`
    TenantId  string `json:"tenantId,omitempty"`
}

