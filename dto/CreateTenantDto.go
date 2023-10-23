package dto


type CreateTenantDto struct{
    Name  string `json:"name"`
    AppIds  []string `json:"appIds"`
    Logo  []string `json:"logo,omitempty"`
    Description  string `json:"description,omitempty"`
    RejectHint  string `json:"rejectHint,omitempty"`
    SourceAppId  string `json:"sourceAppId,omitempty"`
    EnterpriseDomains  []string `json:"enterpriseDomains,omitempty"`
    ExpireTime  string `json:"expireTime,omitempty"`
    MauAmount  int `json:"mauAmount,omitempty"`
    MemberAmount  int `json:"memberAmount,omitempty"`
    AdminAmount  int `json:"adminAmount,omitempty"`
}

