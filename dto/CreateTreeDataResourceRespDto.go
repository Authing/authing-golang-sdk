package dto


type CreateTreeDataResourceRespDto struct{
    ResourceName  string `json:"resourceName"`
    ResourceCode  string `json:"resourceCode"`
    Type  string  `json:"type"`
    Description  string `json:"description,omitempty"`
    Struct  []DataResourceTreeStructs `json:"struct"`
    Actions  []string `json:"actions"`
}

