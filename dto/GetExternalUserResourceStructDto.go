package dto


type GetExternalUserResourceStructDto struct{
    ResourceCode  string `json:"resourceCode"`
    ExternalId  string `json:"externalId"`
    NamespaceCode  string `json:"namespaceCode"`
}

