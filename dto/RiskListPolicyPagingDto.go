package dto


type RiskListPolicyPagingDto struct{
    TotalCount  int `json:"totalCount"`
    List  []RiskListPolicyRespDto `json:"list"`
}

