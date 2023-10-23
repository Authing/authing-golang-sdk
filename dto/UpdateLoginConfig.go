package dto


type UpdateLoginConfig struct{
    TabMethodsSortConfig  ApplicationTabMethodsSortConfigDto `json:"tabMethodsSortConfig"`
    QrCodeSortConfig  ApplicationTabMethodsSortConfigDto `json:"qrCodeSortConfig"`
    SsoPageCustomizationSettings  ISsoPageCustomizationSettingsDto `json:"ssoPageCustomizationSettings"`
    PasswordTabConfig  TabConfigDto `json:"passwordTabConfig"`
    VerifyCodeTabConfig  TabConfigDto `json:"verifyCodeTabConfig"`
    Config  LanguageCoinfigDto `json:"config"`
    EnableCreateTenant  bool `json:"enableCreateTenant"`
    CreateTenantScenes  []string `json:"createTenantScenes"`
    EnableJoinTenant  bool `json:"enableJoinTenant"`
    JoinTenantScenes  []string `json:"joinTenantScenes"`
    EnableVerifyDomain  bool `json:"enableVerifyDomain"`
    VerifyDomainScenes  []string `json:"verifyDomainScenes"`
}

