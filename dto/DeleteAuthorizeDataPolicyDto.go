package dto


type DeleteAuthorizeDataPolicyDto struct{
    TargetType  string  `json:"targetType"`
    TargetIdentifier  string `json:"targetIdentifier"`
    PolicyId  string `json:"policyId"`
}

