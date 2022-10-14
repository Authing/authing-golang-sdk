package dto


type ApplicationDto struct{
    AppId  string `json:"appId"`
    AppIdentifier  string `json:"appIdentifier"`
    AppName  string `json:"appName"`
    AppLogo  string `json:"appLogo"`
    AppDescription  string `json:"appDescription,omitempty"`
    AppType  string  `json:"appType"`
    UserPoolId  string `json:"userPoolId"`
    IsIntegrateApp  bool `json:"isIntegrateApp"`
    DefaultProtocol  string  `json:"defaultProtocol"`
    RedirectUris  []string `json:"redirectUris"`
    LogoutRedirectUris  []string `json:"logoutRedirectUris"`
    InitLoginUri  string `json:"initLoginUri"`
    SsoEnabled  bool `json:"ssoEnabled"`
    SsoEnabledAt  string `json:"ssoEnabledAt,omitempty"`
    LoginConfig  ApplicationLoginConfigDto `json:"loginConfig"`
    RegisterConfig  ApplicationRegisterConfig `json:"registerConfig"`
    BrandingConfig  ApplicationBrandingConfig `json:"brandingConfig"`
    OidcConfig  OIDCConfig `json:"oidcConfig"`
    SamlProviderEnabled  bool `json:"samlProviderEnabled"`
    SamlConfig  SamlIdpConfig `json:"samlConfig,omitempty"`
    OauthProviderEnabled  bool `json:"oauthProviderEnabled"`
    OauthConfig  OauthIdpConfig `json:"oauthConfig,omitempty"`
    CasProviderEnabled  bool `json:"casProviderEnabled"`
    CasConfig  CasIdPConfig `json:"casConfig,omitempty"`
    CustomBrandingEnabled  bool `json:"customBrandingEnabled"`
    CustomSecurityEnabled  bool `json:"customSecurityEnabled"`
    Template  string `json:"template,omitempty"`
}

