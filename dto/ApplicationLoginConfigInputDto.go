package dto


type ApplicationLoginConfigInputDto struct{
    MergeLoginAndRegisterPage  bool `json:"mergeLoginAndRegisterPage,omitempty"`
    EnabledBasicLoginMethods  []string `json:"enabledBasicLoginMethods,omitempty"`
    DefaultLoginMethod  ApplicationDefaultLoginMethodInput `json:"defaultLoginMethod,omitempty"`
    EnabledExtIdpConnIds  []ApplicationEnabledExtIdpConnInputDto `json:"enabledExtIdpConnIds,omitempty"`
    EnabledAllExtIdpConns  bool `json:"enabledAllExtIdpConns,omitempty"`
    ShowAuthorizationPage  bool `json:"showAuthorizationPage"`
}

