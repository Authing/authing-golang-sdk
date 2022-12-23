package dto


type ListSimpleDataPoliciesRespDto struct{
    PolicyId  string `json:"policyId"`
    PolicyName  string `json:"policyName"`
    Description  string `json:"description,omitempty"`
}

