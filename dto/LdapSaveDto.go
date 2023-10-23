package dto


type LdapSaveDto struct{
    LdapDomain  string `json:"ldapDomain"`
    LinkUrl  string `json:"linkUrl,omitempty"`
}

