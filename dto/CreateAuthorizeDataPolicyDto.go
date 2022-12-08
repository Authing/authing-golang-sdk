package dto


type CreateAuthorizeDataPolicyDto struct{
    TargetList  []SubjectDto `json:"targetList"`
    PolicyIds  []string `json:"policyIds"`
}

