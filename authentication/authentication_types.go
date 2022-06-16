package client

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthenticationClientOptions struct {
	/** 应用 ID */
	AppId string

	/** 应用 Secret */
	AppSecret string

	/** 应用对应的用户池域名，例如 pool.authing.cn */
	Domain string

	/** 认证完成后的重定向目标 URL */
	RedirectUri string

	/** 登出完成后的重定向目标 URL, 可选 */
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
	 * 服务端的 JWKS 公钥，用于验证 Token 签名，默认会通过网络请求从服务端的 JWKS 端点自动获取
	 */
	// serverJWKS JWKSObject;
}
type AuthUrlResult struct {
	Url   string
	State string
	Nonce string
}

type AuthURLParams struct {
	RedirectUri string
	State       string
	Nonce       string
	Scope       string
	Forced      bool
}

type CodeToTokenParams struct {
	Code        string
	RedirectUri string
}

// type RefreshTokenParams struct {
// 	grant_type    string //'refresh_token'
// 	client_id     string
// 	client_secret string
// 	refresh_token string
// }

type LoginState struct {
	AccessToken       string `json:"access_token"`
	IdToken           string `json:"id_token"`
	RefreshToken      string `json:"refresh_token"` //可选
	ExpiresIn         uint64 `json:"expires_in"`
	ExpireAt          time.Time
	ParsedIDToken     *IDTokenClaims
	ParsedAccessToken *AccessTokenClaims
}

// type LoginTransaction struct {
// 	state       string
// 	nonce       string
// 	redirectUri string
// }
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
	Sub string `json:"sub,omitempty"`
	UserInfoCommon
}

type TokenCommon struct {
	Sub string `json:"sub,omitempty"`
	Aud string `json:"aud,omitempty"`
	Exp uint64 `json:"exp,omitempty"`
	Iat uint64 `json:"iat,omitempty"`
	Iss string `json:"iss,omitempty"`
	Jti string `json:"jti,omitempty"`
}
type IDTokenExtended struct {
	Nonce  string `json:"nonce,omitempty"`
	AtHash string `json:"at_hash,omitempty"`
	SHash  string `json:"s_hash,omitempty"`
}
type IDToken struct {
	UserInfoCommon //扩展自UserInfo
	TokenCommon
	IDTokenExtended
}

type IDTokenClaims struct {
	UserInfoCommon
	IDTokenExtended
	jwt.StandardClaims
}
type AccessTokenExtended struct {
	Scope string `json:"scope,omitempty"`
}
type AccessToken struct {
	TokenCommon
	AccessTokenExtended
}

type AccessTokenClaims struct {
	jwt.StandardClaims
	AccessTokenExtended
}

type LogoutURLParams struct {
	RedirectUri string // 可选
	IDTokenHint string // 可选
	State       string // 可选
}
