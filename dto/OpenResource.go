package dto


type OpenResource struct{
    ResourceCode  string `json:"resourceCode"`
    ResourceType  string  `json:"resourceType"`
    StrAuthorize  StrAuthorize `json:"strAuthorize,omitempty"`
    ArrAuthorize  ArrayAuthorize `json:"arrAuthorize,omitempty"`
    TreeAuthorize  TreeAuthorize `json:"treeAuthorize,omitempty"`
}

