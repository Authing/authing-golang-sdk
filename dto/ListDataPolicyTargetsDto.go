package dto


type ListDataPolicyTargetsDto struct{
    PolicyId string `json:"policyId,omitempty"`
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
    Query string `json:"query,omitempty"`
    TargetType string `json:"targetType,omitempty"`
}

