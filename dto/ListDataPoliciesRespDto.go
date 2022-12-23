package dto


type ListDataPoliciesRespDto struct{
    PolicyName  string `json:"policyName"`
    Description  string `json:"description,omitempty"`
    ResourceList  []DataResourceSimpleRespDto `json:"resourceList"`
    PolicyId  string `json:"policyId"`
    TargetList  []SubjectRespDto `json:"targetList"`
    UpdatedAt  string `json:"updatedAt"`
}

