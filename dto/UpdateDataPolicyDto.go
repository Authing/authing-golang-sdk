package dto


type UpdateDataPolicyDto struct{
    PolicyId  string `json:"policyId"`
    PolicyName  string `json:"policyName,omitempty"`
    Description  string `json:"description,omitempty"`
    StatementList  []DataStatementPermissionDto `json:"statementList,omitempty"`
}

