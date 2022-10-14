package dto


type ApplicationLoginConfigDto struct{
    MergeLoginAndRegisterPage  bool `json:"mergeLoginAndRegisterPage"`
    EnabledBasicLoginMethods  []string `json:"enabledBasicLoginMethods"`
    DefaultLoginMethod  ApplicationDefaultLoginMethod `json:"defaultLoginMethod"`
    EnabledExtIdpConns  []ApplicationEnabledExtIdpConnDto `json:"enabledExtIdpConns"`
    ShowAuthorizationPage  bool `json:"showAuthorizationPage"`
}

