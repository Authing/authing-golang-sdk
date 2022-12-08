package dto


type DataPoliciesPaginatedRespDto struct{
    TotalCount  int `json:"totalCount"`
    List  []ListDataPoliciesRespDto `json:"list"`
}

