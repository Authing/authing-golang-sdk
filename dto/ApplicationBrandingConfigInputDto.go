package dto


type ApplicationBrandingConfigInputDto struct{
    CustomCSSEnabled  bool `json:"customCSSEnabled,omitempty"`
    CustomCSS  string `json:"customCSS,omitempty"`
    GuardVersion  string  `json:"guardVersion,omitempty"`
    CustomLoadingImage  string `json:"customLoadingImage,omitempty"`
    CustomBackground  string `json:"customBackground,omitempty"`
    ShowChangeLanguageButton  bool `json:"showChangeLanguageButton,omitempty"`
    DefaultLanguage  string  `json:"defaultLanguage,omitempty"`
    ShowForgetPasswordButton  bool `json:"showForgetPasswordButton,omitempty"`
    ShowEnterpriseConnections  bool `json:"showEnterpriseConnections,omitempty"`
    ShowSocialConnections  bool `json:"showSocialConnections,omitempty"`
}

