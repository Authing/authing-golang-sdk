package dto


type TenantDto struct{
    TenantId  string `json:"tenantId"`
    UserPoolId  string `json:"userPoolId"`
    Name  string `json:"name"`
    Description  string `json:"description,omitempty"`
    Logo  []string `json:"logo"`
    RejectHint  string `json:"rejectHint,omitempty"`
    AppIds  []string `json:"appIds"`
    Creator  string `json:"creator"`
    SourceAppId  string `json:"sourceAppId"`
    Source  string `json:"source"`
    Code  string `json:"code"`
    EnterpriseDomains  string `json:"enterpriseDomains"`
}

