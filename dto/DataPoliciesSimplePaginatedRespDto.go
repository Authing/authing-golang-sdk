package dto


type DataPoliciesSimplePaginatedRespDto struct{
    TotalCount  int `json:"totalCount,omitempty"`
    List  []ListSimpleDataPoliciesRespDto `json:"list"`
}

