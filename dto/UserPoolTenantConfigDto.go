package dto


type UserPoolTenantConfigDto struct{
    UserPoolId  string `json:"userPoolId"`
    IsUserPoolAsTenant  bool `json:"isUserPoolAsTenant"`
    EnableSwitchType  string  `json:"enableSwitchType"`
    Css  string `json:"css"`
    CssEnabled  bool `json:"cssEnabled"`
    CustomLoading  string `json:"customLoading"`
    EnableGuardVersionSwitch  bool `json:"enableGuardVersionSwitch"`
    GuardVersion  string `json:"guardVersion"`
    LoadingBackground  string `json:"loadingBackground"`
    EnableCreateTenant  bool `json:"enableCreateTenant"`
    CreateTenantScenes  []string `json:"createTenantScenes"`
    EnableJoinTenant  bool `json:"enableJoinTenant"`
    JoinTenantScenes  []string `json:"joinTenantScenes"`
    EnableVerifyDomain  bool `json:"enableVerifyDomain"`
    VerifyDomainScenes  []string `json:"verifyDomainScenes"`
    SsoPageCustomizationSettings  ISsoPageCustomizationSettingsDto `json:"ssoPageCustomizationSettings"`
    EnableMultipleTenantPortal  bool `json:"enableMultipleTenantPortal"`
    LoginConfig  ApplicationLoginConfigDto `json:"loginConfig"`
}

