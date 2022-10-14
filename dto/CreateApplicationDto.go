package dto


type CreateApplicationDto struct{
    AppName  string `json:"appName"`
    Template  string `json:"template,omitempty"`
    TemplateData  string `json:"templateData,omitempty"`
    AppIdentifier  string `json:"appIdentifier,omitempty"`
    AppLogo  string `json:"appLogo,omitempty"`
    AppDescription  string `json:"appDescription,omitempty"`
    AppType  string  `json:"appType,omitempty"`
    DefaultProtocol  string  `json:"defaultProtocol,omitempty"`
    RedirectUris  []string `json:"redirectUris,omitempty"`
    LogoutRedirectUris  []string `json:"logoutRedirectUris,omitempty"`
    InitLoginUri  string `json:"initLoginUri,omitempty"`
    SsoEnabled  bool `json:"ssoEnabled,omitempty"`
    OidcConfig  OIDCConfig `json:"oidcConfig,omitempty"`
    SamlProviderEnabled  bool `json:"samlProviderEnabled,omitempty"`
    SamlConfig  SamlIdpConfig `json:"samlConfig,omitempty"`
    OauthProviderEnabled  bool `json:"oauthProviderEnabled,omitempty"`
    OauthConfig  OauthIdpConfig `json:"oauthConfig,omitempty"`
    CasProviderEnabled  bool `json:"casProviderEnabled,omitempty"`
    CasConfig  CasIdPConfig `json:"casConfig,omitempty"`
    LoginConfig  ApplicationLoginConfigInputDto `json:"loginConfig,omitempty"`
    RegisterConfig  ApplicationRegisterConfigInputDto `json:"registerConfig,omitempty"`
    BrandingConfig  ApplicationBrandingConfigInputDto `json:"brandingConfig,omitempty"`
}

