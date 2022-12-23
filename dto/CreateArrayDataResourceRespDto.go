package dto


type CreateArrayDataResourceRespDto struct{
    ResourceName  string `json:"resourceName"`
    ResourceCode  string `json:"resourceCode"`
    Type  string  `json:"type"`
    Description  string `json:"description,omitempty"`
    Struct  []string `json:"struct"`
    Actions  []string `json:"actions"`
}

