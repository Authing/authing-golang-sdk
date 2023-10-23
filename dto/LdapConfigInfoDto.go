package dto


type LdapConfigInfoDto struct{
    Id  string `json:"id,omitempty"`
    Enabled  int `json:"enabled"`
    UserPoolId  string `json:"userPoolId"`
    LinkUrl  string `json:"linkUrl"`
    LdapsLinkUrl  string `json:"ldapsLinkUrl"`
    EnablePrivatization  int `json:"enablePrivatization"`
    BindDn  string `json:"bindDn"`
    LdapDomain  string `json:"ldapDomain"`
    EnableSsl  int `json:"enableSsl,omitempty"`
    BaseDn  string `json:"baseDn"`
    BindPwd  string `json:"bindPwd"`
    VisibleOrgNodes  interface{} `json:"visibleOrgNodes"`
    VisibleFields  interface{} `json:"visibleFields"`
    UdfMapping  interface{} `json:"udfMapping"`
}

