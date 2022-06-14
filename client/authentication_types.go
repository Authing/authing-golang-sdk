package client

import "time"

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

	/**
	 * 存储认证上下文的 Cookie 名称
	 */
	CookieKey string
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

type OIDCTokenResponse struct {
	access_token  string
	id_token      string
	expires_in    uint64
	refresh_token string
	token_type    string
}

type NativeTokenParams struct { //TODO 允许扩展
	grant_type    string
	client_id     string
	client_secret string
	scope         string
}

type CodeToTokenParams struct {
	Code        string
	RedirectUri string
}

type RefreshTokenParams struct {
	grant_type    string //'refresh_token'
	client_id     string
	client_secret string
	refresh_token string
}

type LoginState struct {
	AccessToken       string `json:"access_token"`
	IdToken           string `json:"id_token"`
	RefreshToken      string `json:"refresh_token"` //可选
	expiresIn         uint64 `json:"expires_in"`
	ExpireAt          time.Time
	ParsedIDToken     *IDToken
	ParsedAccessToken *AccessToken
}

type LoginTransaction struct {
	state       string
	nonce       string
	redirectUri string
}

type UserInfo struct { // TODO 允许扩展
	Sub               string
	Name              string
	Nickname          string
	GivenName         string
	FamilyName        string
	Birthdate         string
	Gender            string //'M' | 'F' | 'U'
	Picture           string
	UpdatedAt         string
	Zoneinfo          string
	PreferredUsername string
	Locale            string
}

type IDToken struct {
	UserInfo //扩展自UserInfo
	Sub      string
	Aud      string
	Exp      uint64
	Iat      uint64
	Iss      string
	Nonce    string
	AtHash   string
	SHash    string
}

type AccessToken struct {
	Jti   string
	Sub   string
	Iat   uint64
	Exp   uint64
	Scope uint64
	Iss   string
	Aud   string
}

//   type IDToken struct {// TODO 允许扩展
// 	// Token 字段
// 	sub string
// 	aud string
// 	exp uint64
// 	iat uint64
// 	iss string
// 	nonce string
// 	at_hash string
// 	s_hash string

// 	// 用户信息字段
// 	name string
// 	nickname string
// 	given_name string
// 	family_name string
// 	birthdate string
// 	gender string//'M' | 'F' | 'U'
// 	picture string
// 	updated_at string
// 	zoneinfo string
// 	preferred_username string
// 	locale string
//   }

type LogoutURLParams struct {
	post_logout_redirect_uri string // 可选
	id_token_hint            string // 可选
	state                    string // 可选
}
