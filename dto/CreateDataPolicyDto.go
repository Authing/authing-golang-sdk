package dto


type CreateDataPolicyDto struct{
    StatementList  []DataStatementPermissionDto `json:"statementList"`
    PolicyName  string `json:"policyName"`
    Description  string `json:"description,omitempty"`
}

