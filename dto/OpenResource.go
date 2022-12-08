package dto


type OpenResource struct{
    ResourceCode  string `json:"resourceCode"`
    Authorize  (TreeAuthorize | ArrayAuthorize | StrAuthorize) `json:"authorize"`
}

