package dto


type UpdateOrganizationStatusReqDto struct{
    RootNodeId  string `json:"rootNodeId"`
    Status  string `json:"status,omitempty"`
}

