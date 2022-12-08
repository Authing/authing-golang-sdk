package dto


type UpdateDataPolicyDto struct{
    PolicyName  string `json:"policyName"`
    PolicyId  string `json:"policyId"`
    Description  string `json:"description,omitempty"`
    StatementList  []DataStatementPermissionDto `json:"statementList,omitempty"`
}

