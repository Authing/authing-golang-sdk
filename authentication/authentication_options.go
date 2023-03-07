package authentication

import (
	"github.com/dgrijalva/jwt-go"
)

type AuthenticationClientOptions struct {
	/**
	应用 ID
	*/
	AppId string

	/**
	应用 Secret
	*/
	AppSecret string

	/**
	应用域名，例如 pool.authing.cn
	*/
	AppHost string

	/**
	用户的 Access Token
	*/
	AccessToken string

	/**
	认证完成后的重定向目标 URL
	*/
	RedirectUri string

	/**
	登出完成后的重定向目标 URL, 可选
	*/
	LogoutRedirectUri string

	/**
	 * @optional
	 * 应用侧向 Authing 请求的权限，以空格分隔，默认为 'openid profile'
	 *
	 * 成功获取的权限会出现在 Access Token 的 scope 字段中
	 *
	 * 一些示例：
	 * - openid OIDC 标准规定的权限，必须包含
	 * - profile 获取用户的基本身份信息
	 * - offline_access 获取用户的 Refresh Token，可用于调用 refreshLoginState 刷新用户的登录态
	 */
	Scope string

	/**
	获取 token 端点认证方式
	*/
	TokenEndPointAuthMethod TokenAuthMethodEnum

	/**
	检测 token 端点认证方式
	*/
	IntrospectionEndPointAuthMethod TokenAuthMethodEnum

	/**
	撤销 token 端点认证方式
	*/
	RevocationEndPointAuthMethod TokenAuthMethodEnum

	/**
	协议类型，默认为 oidc
	*/
	Protocol ProtocolEnum

	/**
	请求超时时间
	*/
	Timeout int

	/**
	是否跳过 HTTPS 证书检测，默认为 false；如果是私有化部署的场景且证书不被信任，可以设置为 true
	*/
	InsecureSkipVerify bool
	/**
	 * 订阅事件 WebSocket 地址
	 */
	WssHost string
}

type AuthUrlResult struct {
	Url   string
	State string
	Nonce string
}

type OIDCAuthURLParams struct {
	RedirectUri  string
	State        string
	Nonce        string
	Scope        string
	ResponseType string
	ResponseMode string
	Forced       bool
}

type OAuth2AuthURLParams struct {
	RedirectUri string
	State       string
	/**
	默认为 openid profile email phone address
	*/
	Scope        string
	ResponseType string
}

type CodeToTokenParams struct {
	Code        string
	RedirectUri string
	Nonce       string
}

type OIDCTokenResponse struct {
	AccessToken      string `json:"access_token,omitempty"`
	IDToken          string `json:"id_token,omitempty"`
	RefreshToken     string `json:"refresh_token,omitempty"` //可选
	ExpiresIn        uint64 `json:"expires_in,omitempty"`
	TokenType        string `json:"token_type,omitempty"`
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type UserInfoCommon struct {
	Name              string `json:"name,omitempty"`
	Nickname          string `json:"nickname,omitempty"`
	GivenName         string `json:"given_name,omitempty"`
	FamilyName        string `json:"family_name,omitempty"`
	Birthdate         string `json:"birthdate,omitempty"`
	Gender            string `json:"gender,omitempty"` //'M' | 'F' | 'U'
	Picture           string `json:"picture,omitempty"`
	UpdatedAt         string `json:"updated_at,omitempty"`
	Zoneinfo          string `json:"zoneinfo,omitempty"`
	PreferredUsername string `json:"preferred_username,omitempty"`
	Locale            string `json:"locale,omitempty"`
}
type UserInfo struct { // TODO 允许扩展
	Subject string `json:"sub,omitempty"` // 用户 ID
	UserInfoCommon
}

type IDTokenExtended struct {
	Nonce  string `json:"nonce,omitempty"`
	AtHash string `json:"at_hash,omitempty"`
	SHash  string `json:"s_hash,omitempty"`
}

type IDTokenClaims struct {
	UserInfoCommon
	IDTokenExtended
	jwt.StandardClaims
}
type AccessTokenExtended struct {
	Scope string `json:"scope,omitempty"`
}

type AccessTokenClaims struct {
	jwt.StandardClaims
	AccessTokenExtended
}

type BuildLogoutURLParams struct {
	PostLogoutRedirectUri string // 可选
	IDTokenHint           string // 可选
	State                 string // 可选
}

type ClientCredentialInput struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

type GetAccessTokenByClientCredentialsRequest struct {
	Scope                 string                 `json:"scope"`
	ClientCredentialInput *ClientCredentialInput `json:"client_credential_input"`
}

type TokenAuthMethodEnum string

const (
	ClientSecretPost  = "client_secret_post"
	ClientSecretBasic = "client_secret_basic"
	None              = "none"
)

type ProtocolEnum string

const (
	OAUTH ProtocolEnum = "oauth"
	OIDC  ProtocolEnum = "oidc"
	CAS   ProtocolEnum = "cas"
	SAML  ProtocolEnum = "saml"
)
