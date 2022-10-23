package dto


type UpdateTenantDto struct{
    SsoPageCustomizationSettings  string `json:"ssoPageCustomizationSettings"`
    ExtendsFields  string `json:"extendsFields"`
    RegisterTabs  string `json:"registerTabs"`
    LoginTabs  string `json:"loginTabs"`
    PasswordTabConfig  string `json:"passwordTabConfig"`
    DefaultRegisterTab  string `json:"defaultRegisterTab"`
    DefaultLoginTab  string `json:"defaultLoginTab"`
    Css  string `json:"css"`
    AppIds  []string `json:"appIds"`
    Name  string `json:"name"`
    TenantId  string `json:"tenantId"`
}

