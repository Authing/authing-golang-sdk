package dto

type SignInOptionsDto struct {
	Scope               string      `json:"scope,omitempty"`
	ClientIp            string      `json:"clientIp,omitempty"`
	Context             interface{} `json:"context,omitempty"`
	TenantId            string      `json:"tenantId,omitempty"`
	CustomData          interface{} `json:"customData,omitempty"`
	AutoRegister        bool        `json:"autoRegister"`
	CaptchaCode         string      `json:"captchaCode,omitempty"`
	PasswordEncryptType string      `json:"passwordEncryptType,omitempty"`
}
