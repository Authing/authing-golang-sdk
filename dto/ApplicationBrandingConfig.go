package dto


type ApplicationBrandingConfig struct{
    CustomCSSEnabled  bool `json:"customCSSEnabled"`
    CustomCSS  string `json:"customCSS,omitempty"`
    GuardVersion  string  `json:"guardVersion"`
    CustomLoadingImage  string `json:"customLoadingImage,omitempty"`
    CustomBackground  string `json:"customBackground,omitempty"`
    ShowChangeLanguageButton  bool `json:"showChangeLanguageButton"`
    DefaultLanguage  string  `json:"defaultLanguage"`
    ShowForgetPasswordButton  bool `json:"showForgetPasswordButton"`
    ShowEnterpriseConnections  bool `json:"showEnterpriseConnections"`
    ShowSocialConnections  bool `json:"showSocialConnections"`
    ShowAgreement  bool `json:"showAgreement"`
    Agreements  []ApplicationAgreementDto `json:"agreements"`
}

