package client

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
	url   string
	state string
	nonce string
}

type AuthURLParams struct {
	redirect_uri  string
	response_type string
	response_mode string
	client_id     string
	state         string
	nonce         string
	scope         string
	prompt        string
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
	grant_type    string //'authorization_code'
	client_id     string
	client_secret string
	code          string
	redirect_uri  string
}

type RefreshTokenParams struct {
	grant_type    string //'refresh_token'
	client_id     string
	client_secret string
	refresh_token string
}

type LoginState struct {
	accessToken       string
	idToken           string
	refreshToken      string //可选
	expireAt          uint64
	parsedIDToken     IDToken
	parsedAccessToken AccessToken
}

type LoginTransaction struct {
	state       string
	nonce       string
	redirectUri string
}

type UserInfo struct { // TODO 允许扩展
	sub                string
	name               string
	nickname           string
	given_name         string
	family_name        string
	birthdate          string
	gender             string //'M' | 'F' | 'U'
	picture            string
	updated_at         string
	zoneinfo           string
	preferred_username string
	locale             string
}

type IDToken struct {
	UserInfo //扩展自UserInfo
	sub      string
	aud      string
	exp      uint64
	iat      uint64
	iss      string
	nonce    string
	at_hash  string
	s_hash   string
}

type AccessToken struct {
	jti   string
	sub   string
	iat   uint64
	exp   uint64
	scope uint64
	iss   string
	aud   string
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
