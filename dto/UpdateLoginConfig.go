package dto

type UpdateLoginConfig struct {
	SsoPageCustomizationSettings ISsoPageCustomizationSettingsDto `json:"ssoPageCustomizationSettings"`
	PasswordTabConfig            TabConfigDto                     `json:"passwordTabConfig"`
	VerifyCodeTabConfig          TabConfigDto                     `json:"verifyCodeTabConfig"`
	Config                       LanguageCoinfigDto               `json:"config"`
}
