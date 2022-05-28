package dto


type AuthorizedResourceDto struct{
    ResourceCode  string `json:"resourceCode"`
    ResourceType  string  `json:"resourceType,omitempty"`
    Actions  []string `json:"actions,omitempty"`
    ApiIdentifier  string `json:"apiIdentifier,omitempty"`
}

