package dto


type ListDataPolicySubjectPageDto struct{
    TotalCount  int `json:"totalCount,omitempty"`
    List  []DataSubjectRespDto `json:"list"`
}

