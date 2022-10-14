package authentication

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Authing/authing-golang-sdk/dto"

	// "net/http"
	"strings"
	"time"

	"github.com/Authing/authing-golang-sdk/constant"
	"github.com/Authing/authing-golang-sdk/util"

	keyfunc "github.com/MicahParks/compatibility-keyfunc"
	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
)

const RandStringLen = 16
const ALG_HS256 = "HS256"
const JWK_PATH = "/oidc/.well-known/jwks.json"

type Client struct {
	appId                           string
	appSecret                       string
	appHost                         string
	protocol                        ProtocolEnum
	redirectUri                     string
	logoutRedirectUri               string
	scope                           string
	tokenEndPointAuthMethod         TokenAuthMethodEnum
	introspectionEndPointAuthMethod TokenAuthMethodEnum
	revocationEndPointAuthMethod    TokenAuthMethodEnum
	timeout                         int8
	rejectUnauthorized              bool
	jwks                            *keyfunc.JWKS
}

var commonHeaders = map[string]string{
	"x-authing-request-from": constant.SdkName,
	"x-authing-sdk-version":  constant.SdkVersion,
}

func NewClient(options *AuthenticationClientOptions) (*Client, error) {
	if options.AppId == "" {
		return nil, errors.New("AppId 不能为空")
	}
	if options.AppSecret == "" {
		return nil, errors.New("AppSecret 不能为空")
	}
	if options.AppHost == "" {
		return nil, errors.New("AppHost 不能为空")
	}
	if options.RedirectUri == "" {
		return nil, errors.New("RedirectUri 不能为空")
	}
	client := &Client{
		appId:                           options.AppId,
		appSecret:                       options.AppSecret,
		appHost:                         options.AppHost,
		protocol:                        options.Protocol,
		redirectUri:                     options.RedirectUri,
		logoutRedirectUri:               options.LogoutRedirectUri,
		scope:                           options.Scope,
		tokenEndPointAuthMethod:         options.TokenEndPointAuthMethod,
		introspectionEndPointAuthMethod: options.IntrospectionEndPointAuthMethod,
		revocationEndPointAuthMethod:    options.RevocationEndPointAuthMethod,
		timeout:                         options.Timeout,
		rejectUnauthorized:              options.RejectUnauthorized,
	}
	if options.Scope == "" {
		client.scope = constant.DefaultScope
	}
	if options.Protocol == "" {
		client.protocol = OIDC
	}

	jwksURL := client.getUrl(JWK_PATH)
	jwks, err := keyfunc.Get(jwksURL, keyfunc.Options{})
	if err != nil {
		return nil, fmt.Errorf("获取 jwk 密钥失败: %w", err)
	}
	client.jwks = jwks

	return client, nil
}

func (client *Client) getUrl(path string) string {
	return client.appHost + path
}

func (client *Client) BuildAuthorizeUrlByOidc(params *OIDCAuthURLParams) (AuthUrlResult, error) {
	if params == nil {
		params = &OIDCAuthURLParams{}
	}

	// scope
	scope := params.Scope
	if scope == "" {
		scope = client.scope
	}

	// state 和 nonce
	state := params.State
	nonce := params.Nonce
	if state == "" {
		state = util.RandStringImpr(RandStringLen)
	}
	if nonce == "" {
		nonce = util.RandStringImpr(RandStringLen)
	}
	redirectUri := params.RedirectUri
	if redirectUri == "" {
		redirectUri = client.redirectUri
	}
	responseMode := params.ResponseMode
	if responseMode == "" {
		responseMode = "query"
	}

	dataMap := map[string]interface{}{
		"redirect_uri":  redirectUri,
		"client_id":     client.appId,
		"response_mode": "query",
		"response_type": "code",
		"scope":         scope,
		"nonce":         nonce,
		"state":         state,
	}

	if params.Forced {
		dataMap["prompt"] = "login"
	} else {
		arr := strings.Split(scope, " ")

		for _, value := range arr {
			if value == "offline_access" {
				dataMap["prompt"] = "consent"
				break
			}
		}
	}
	return AuthUrlResult{
		State: state,
		Nonce: nonce,
		Url:   client.getUrl("/oidc/auth?") + util.GenQueryString(dataMap),
	}, nil
}

func (client *Client) BuildAuthorizeUrlByOauth(params *OAuth2AuthURLParams) (string, error) {
	dataMap := map[string]string{
		"client_id":     util.GetValueOrDefault(client.appId),
		"scope":         util.GetValueOrDefault(params.Scope, "openid profile email phone address"),
		"state":         util.GetValueOrDefault(params.State, util.RandomString(12)),
		"response_type": util.GetValueOrDefault(params.ResponseType, "code"),
		"redirect_uri":  util.GetValueOrDefault(params.RedirectUri, client.redirectUri),
	}
	return client.getUrl("/oauth/auth?") + util.GetQueryString(dataMap), nil
}

func (client *Client) BuildAuthorizeUrlBySaml() string {
	return fmt.Sprintf("%s/api/v2/saml-idp/%s", client.appHost, client.appId)
}

func (client *Client) BuildAuthorizeUrlByCas(service *string) string {
	if service != nil {
		return fmt.Sprintf("%s/cas-idp/%s?service=%s", client.appHost, client.appId, *service)
	} else {
		return fmt.Sprintf("%s/cas-idp/%s?service", client.appHost, client.appId)
	}
}

/*
*
检验 CAS 1.0 Ticket 合法性
*/
func (c *Client) ValidateTicketV1(ticket, service string) (*struct {
	Valid    bool   `json:"code"`
	Message  string `json:"message"`
	Username string `json:"username"`
}, error) {
	url := fmt.Sprintf("%s/cas-idp/%s/validate?service=%s&ticket=%s", c.appHost, c.appId, service, ticket)
	res, err := util.SendRequest(&util.RequestOption{
		Url:     url,
		Method:  fasthttp.MethodGet,
		Headers: c.getReqHeaders(nil),
	})
	if err != nil {
		return nil, err
	}
	sps := strings.Split(string(res.Body), "\n")
	var username, message string

	valid := (sps[0] == "yes")
	username = sps[1]
	if !valid {
		message = "ticket is not valid"
	}
	resp := &struct {
		Valid    bool   `json:"code"`
		Message  string `json:"message"`
		Username string `json:"username"`
	}{
		Valid:    valid,
		Username: username,
		Message:  message,
	}

	return resp, nil
}

// 通过远端服务验证票据合法性
func (c *Client) ValidateTicketV2(ticket, service string, format string) (*struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}, error) {
	url := fmt.Sprintf("%s/cas-idp/%s/serviceValidate", c.appHost, c.appId)
	res, err := util.SendRequest(&util.RequestOption{
		Url:     url,
		Method:  fasthttp.MethodGet,
		Headers: c.getReqHeaders(nil),
		ReqDto: map[string]interface{}{
			"service": service,
			"ticket":  ticket,
			"format":  format,
		},
	})
	if err != nil {
		return nil, err
	}
	resp := &struct {
		Code    int64       `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{}
	json.Unmarshal(res.Body, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return resp, nil
}

// GetAccessTokenByCode 使用 code 换取 accessToken
func (c *Client) GetAccessTokenByCode(code string) (string, error) {
	if c.appId == "" {
		return constant.StringEmpty, errors.New("请在初始化 AuthenticationClient 时传入 appId")
	}
	if c.appSecret == "" && c.tokenEndPointAuthMethod != constant.None {
		return constant.StringEmpty, errors.New("请在初始化 AuthenticationClient 时传入 appSecret")
	}
	url := c.appHost + "/oidc/token"
	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	body := map[string]string{
		"client_id":     c.appId,
		"client_secret": c.appSecret,
		"grant_type":    "authorization_code",
		"code":          code,
		"redirect_uri":  c.redirectUri,
	}

	switch c.tokenEndPointAuthMethod {
	case ClientSecretPost:
		body["client_id"] = c.appId
		body["client_secret"] = c.appSecret
	case ClientSecretBasic:
		base64String := "Basic " + base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.appId, c.appSecret)))
		header["Authorization"] = base64String
	default:
		body["client_id"] = c.appId
	}
	resp, err := util.SendRequest(&util.RequestOption{
		Url:     url,
		Method:  fasthttp.MethodPost,
		Headers: c.getReqHeaders(header),
		ReqDto:  body,
	})
	return string(resp.Body), err
}

// GetAccessTokenByClientCredentials
// Client Credentials 模式获取 Access Token
func (c *Client) GetAccessTokenByClientCredentials(req GetAccessTokenByClientCredentialsRequest) (string, error) {
	if req.Scope == constant.StringEmpty {
		return constant.StringEmpty, errors.New("请传入 scope 参数，请看文档：https://docs.authing.cn/v2/guides/authorization/m2m-authz.html")
	}
	if req.ClientCredentialInput == nil {
		return constant.StringEmpty, errors.New("请在调用本方法时传入 ClientCredentialInput 参数，请看文档：https://docs.authing.cn/v2/guides/authorization/m2m-authz.html")
	}

	url := c.appHost + "/oidc/token"

	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	body := map[string]string{
		"client_id":     req.ClientCredentialInput.AccessKey,
		"client_secret": req.ClientCredentialInput.SecretKey,
		"grant_type":    "client_credentials",
		"scope":         req.Scope,
	}

	resp, err := util.SendRequest(&util.RequestOption{
		Url:     url,
		Method:  fasthttp.MethodPost,
		Headers: c.getReqHeaders(header),
		ReqDto:  body,
	})
	return string(resp.Body), err
}

// GetNewAccessTokenByRefreshToken
//
//	使用 Refresh token 获取新的 Access token
func (c *Client) GetNewAccessTokenByRefreshToken(refreshToken string) (string, error) {
	if c.protocol != OIDC && c.protocol != OAUTH {
		return constant.StringEmpty, errors.New("初始化 AuthenticationClient 时传入的 protocol 参数必须为 ProtocolEnum.OAUTH 或 ProtocolEnum.OIDC，请检查参数")
	}
	if c.appSecret == "" && c.tokenEndPointAuthMethod != constant.None {
		return constant.StringEmpty, errors.New("请在初始化 AuthenticationClient 时传入 Secret")
	}

	url := c.appHost + fmt.Sprintf("/%s/token", c.protocol)

	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	body := map[string]string{
		"client_id":     c.appId,
		"client_secret": c.appSecret,
		"grant_type":    "refresh_token",
		"refresh_token": refreshToken,
	}

	switch c.tokenEndPointAuthMethod {
	case ClientSecretPost:
		body["client_id"] = c.appId
		body["client_secret"] = c.appSecret
	case ClientSecretBasic:
		base64String := "Basic " + base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.appId, c.appSecret)))
		header["Authorization"] = base64String
	default:
		body["client_id"] = c.appId
	}
	resp, err := util.SendRequest(&util.RequestOption{
		Url:     url,
		Method:  fasthttp.MethodPost,
		Headers: c.getReqHeaders(header),
		ReqDto:  body,
	})
	return string(resp.Body), err
}

func (c *Client) IntrospectToken(token string) (string, error) {
	url := c.appHost + fmt.Sprintf("/%s/token/introspection", c.protocol)
	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	body := map[string]string{
		"token": token,
	}
	switch c.tokenEndPointAuthMethod {
	case ClientSecretPost:
		body["client_id"] = c.appId
		body["client_secret"] = c.appSecret
	case ClientSecretBasic:
		base64String := "Basic " + base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.appId, c.appSecret)))
		header["Authorization"] = base64String
	default:
		body["client_id"] = c.appId
	}
	resp, err := util.SendRequest(&util.RequestOption{
		Url:     url,
		Method:  fasthttp.MethodPost,
		Headers: c.getReqHeaders(header),
		ReqDto:  body,
	})
	return string(resp.Body), err
}

// RevokeToken
// 撤回 Access token 或 Refresh token
func (c *Client) RevokeToken(token string) (string, error) {
	if c.protocol != OIDC && c.protocol != OAUTH {
		return constant.StringEmpty, errors.New("初始化 AuthenticationClient 时传入的 protocol 参数必须为 ProtocolEnum.OAUTH 或 ProtocolEnum.OIDC，请检查参数")
	}
	if c.appSecret == "" && c.tokenEndPointAuthMethod != constant.None {
		return constant.StringEmpty, errors.New("请在初始化 AuthenticationClient 时传入 Secret")
	}

	url := c.appHost + fmt.Sprintf("/%s/token/revocation", c.protocol)

	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	body := map[string]string{
		"client_id": c.appId,
		"token":     token,
	}

	switch c.tokenEndPointAuthMethod {
	case ClientSecretPost:
		body["client_id"] = c.appId
		body["client_secret"] = c.appSecret
	case ClientSecretBasic:
		base64String := "Basic " + base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.appId, c.appSecret)))
		header["Authorization"] = base64String
	default:
		body["client_id"] = c.appId
	}
	resp, err := util.SendRequest(&util.RequestOption{
		Url:     url,
		Method:  fasthttp.MethodPost,
		Headers: c.getReqHeaders(header),
		ReqDto:  body,
	})
	return string(resp.Body), err
}

// BuildLogoutUrl
// 拼接登出 URL
func (c *Client) BuildLogoutUrlNew(redirectUri, idToken, state *string) string {
	var url string
	if c.protocol == CAS {
		if redirectUri != nil {
			url = fmt.Sprintf("%s/cas-idp/logout?url=%s", c.appHost, *redirectUri)
		} else {
			url = fmt.Sprintf("%s/cas-idp/logout", c.appHost)
		}
	} else if c.protocol == OIDC {
		if redirectUri != nil {
			url = fmt.Sprintf("%s/oidc/session/end?id_token_hint=%s&post_logout_redirect_uri=%s", c.appHost, *idToken, *redirectUri)
		} else {
			url = fmt.Sprintf("%s/oidc/session/end", c.appHost)
		}
	} else {
		if redirectUri != nil {
			url = fmt.Sprintf("%s/login/profile/logout?redirect_uri=%s", c.appHost, *redirectUri)
		} else {
			url = fmt.Sprintf("%s/login/profile/logout", c.appHost)
		}
	}
	return url
}

func (client *Client) GetLoginStateByAuthCode(params *CodeToTokenParams) (*LoginState, error) {
	if params == nil {
		params = &CodeToTokenParams{}
	}
	if params.Code == "" {
		return nil, errors.New("code 参数不能为空")
	}
	redirectUrl := params.RedirectUri
	if redirectUrl == "" {
		redirectUrl = client.redirectUri
	}
	data := map[string]interface{}{
		"code":          params.Code,
		"client_id":     client.appId,
		"client_secret": client.appSecret,
		"grant_type":    "authorization_code",
		"redirect_uri":  redirectUrl,
	}
	res, err := util.SendRequest(&util.RequestOption{
		Url:     client.getUrl("/oidc/token"),
		Method:  fasthttp.MethodPost,
		ReqDto:  data,
		Headers: client.getReqHeaders(nil),
	})
	if err != nil {
		return nil, fmt.Errorf("根据code获取token时失败: %w", err)
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("根据code获取token失败[%d]:%s", res.StatusCode, res.Body)
	}
	var loginState *LoginState
	loginState, err = client.buildLoginState(res.Body)
	if err != nil {
		return nil, err
	}
	if params.Nonce != "" && loginState.ParsedIDToken.Nonce != params.Nonce {
		return nil, errors.New("nonce 参数不匹配")
	}
	return loginState, nil
}

func (client *Client) getKeyCommon(token *jwt.Token) (interface{}, error) {
	alg, ok := token.Header["alg"].(string)
	if !ok {
		return nil, fmt.Errorf("算法字段非法 %v", token.Header["alg"])
	}
	if alg == ALG_HS256 {
		return []byte(client.appSecret), nil
	}
	return client.jwks.KeyfuncLegacy(token)
}

func (client *Client) getKey4IdToken(token *jwt.Token) (interface{}, error) {
	claims := token.Claims.(*IDTokenClaims)
	claims.IssuedAt = 0

	return client.getKeyCommon(token)
}

func (client *Client) getKey4AccessToken(token *jwt.Token) (interface{}, error) {
	claims := token.Claims.(*AccessTokenClaims)
	claims.IssuedAt = 0

	return client.getKeyCommon(token)
}

func (client *Client) ParsedIDToken(tokenStr string) (*IDTokenClaims, error) {
	tokenJwt, err := jwt.ParseWithClaims(tokenStr, &IDTokenClaims{}, client.getKey4IdToken)
	if err != nil {
		return nil, fmt.Errorf("解析id token失败: %w", err)
	}
	claims, ok := tokenJwt.Claims.(*IDTokenClaims)
	if ok && tokenJwt.Valid {
		return claims, nil
	}
	// TODO return error
	return nil, errors.New("id token非法")
}

func (client *Client) ParsedAccessToken(tokenStr string) (*AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &AccessTokenClaims{}, client.getKey4AccessToken)
	if err != nil {
		return nil, fmt.Errorf("解析 access token失败: %w", err)
	}
	claims, ok := token.Claims.(*AccessTokenClaims)
	if ok && token.Valid {
		return claims, nil
	}
	// TODO return error
	return nil, errors.New("access token非法")
}

func (client *Client) buildLoginState(bytes []byte) (*LoginState, error) {
	var loginState LoginState
	err := json.Unmarshal(bytes, &loginState)
	if err != nil {
		return nil, fmt.Errorf("返回的 token 数据无法解析: %w", err)
	}
	var second time.Duration = time.Duration(loginState.ExpiresIn) * time.Second
	var date time.Time = time.Now()
	loginState.ExpireAt = date.Add(second)
	loginState.ParsedIDToken, err = client.ParsedIDToken(loginState.IDToken)
	if err != nil {
		return nil, err
	}
	loginState.ParsedAccessToken, err = client.ParsedAccessToken(loginState.AccessToken)
	if err != nil {
		return nil, err
	}
	return &loginState, nil
}
func (client *Client) getReqHeaders(customHeaders map[string]string) map[string]string {
	newHeaders := make(map[string]string)
	for key, value := range commonHeaders {
		newHeaders[key] = value
	}
	for key, value := range customHeaders {
		newHeaders[key] = value
	}

	return newHeaders
}
func (client *Client) GetUserInfo(accessToken string) (*UserInfo, error) {
	res, err := util.SendRequest(&util.RequestOption{
		Method: fasthttp.MethodPost,
		Url:    client.getUrl("/oidc/me"),
		Headers: client.getReqHeaders(map[string]string{
			"Authorization": `Bearer ` + accessToken,
		}),
	})
	if err != nil {
		return nil, fmt.Errorf("根据 access token 获取用户信息时失败: %w", err)
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("根据 access token 获取用户信息失败[%d]:%s", res.StatusCode, res.Body)
	}
	var userInfo UserInfo
	err = json.Unmarshal(res.Body, &userInfo)
	if err != nil {
		return nil, fmt.Errorf("无法解析用户信息:%w", err)
	}
	return &userInfo, nil
}

func (client *Client) RefreshLoginState(refreshToken string) (*LoginState, error) {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/oidc/token"),
		Headers: client.getReqHeaders(nil),
		ReqDto: map[string]interface{}{
			"client_id":     client.appId,
			"client_secret": client.appSecret,
			"refresh_token": refreshToken,
			"grant_type":    "refresh_token",
		},
	})
	if err != nil {
		return nil, fmt.Errorf("刷新 token 时失败:%w", err)
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("刷新 token 失败 [%d] %s", res.StatusCode, res.Body)
	}
	return client.buildLoginState(res.Body)
}

func (client *Client) BuildLogoutUrl(params *LogoutURLParams) (string, error) {
	if params == nil {
		params = &LogoutURLParams{}
	}
	idToken := params.IDTokenHint
	redirectUri := params.PostLogoutRedirectUri
	if redirectUri == "" {
		redirectUri = client.logoutRedirectUri
	}
	if redirectUri != "" && idToken == "" {
		return "", errors.New("指定 logout redirect uri 时，必须同时指定 id token 参数")
	}

	data := map[string]interface{}{}
	if redirectUri != "" {
		data["post_logout_redirect_uri"] = redirectUri
	}
	if idToken != "" {
		data["id_token_hint"] = params.IDTokenHint
	}

	if params.State != "" {
		data["state"] = params.State
	}
	queryString := util.GenQueryString(data)
	if queryString != "" {
		queryString = "?" + queryString
	}
	return client.getUrl("/oidc/session/end") + queryString, nil
}

func (client *Client) SignInByUsernamePassword(username string, password string, options dto.SignInOptionsDto) *dto.LoginTokenRespDto {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/signin"),
		Headers: client.getReqHeaders(nil),
		ReqDto: &dto.SigninByCredentialsDto{
			Connection: "PASSWORD",
			PasswordPayload: dto.AuthenticateByPasswordDto{
				Password: password,
				Username: username,
			},
			Options: options,
		},
	})
	var response dto.LoginTokenRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

func (client *Client) SignInByEmailPassword(email string, password string, options dto.SignInOptionsDto) *dto.LoginTokenRespDto {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/signin"),
		Headers: client.getReqHeaders(nil),
		ReqDto: &dto.SigninByCredentialsDto{
			Connection: "PASSWORD",
			PasswordPayload: dto.AuthenticateByPasswordDto{
				Password: password,
				Email:    email,
			},
			Options: options,
		},
	})
	var response dto.LoginTokenRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

func (client *Client) SignInByPhonePassword(phone string, password string, options dto.SignInOptionsDto) *dto.LoginTokenRespDto {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/signin"),
		Headers: client.getReqHeaders(nil),
		ReqDto: &dto.SigninByCredentialsDto{
			Connection: "PASSWORD",
			PasswordPayload: dto.AuthenticateByPasswordDto{
				Password: password,
				Phone:    phone,
			},
			Options: options,
		},
	})
	var response dto.LoginTokenRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

func (client *Client) SignInByAccountPassword(account string, password string, options dto.SignInOptionsDto) *dto.LoginTokenRespDto {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/signin"),
		Headers: client.getReqHeaders(nil),
		ReqDto: &dto.SigninByCredentialsDto{
			Connection: "PASSWORD",
			PasswordPayload: dto.AuthenticateByPasswordDto{
				Password: password,
				Account:  account,
			},
			Options: options,
		},
	})
	var response dto.LoginTokenRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

func (client *Client) SignInByPhonePassCord(phone string, passCode string, phoneCountryCode string, options dto.SignInOptionsDto) *dto.LoginTokenRespDto {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/signin"),
		Headers: client.getReqHeaders(nil),
		ReqDto: &dto.SigninByCredentialsDto{
			Connection: "PASSCODE",
			PassCodePayload: dto.AuthenticateByPassCodeDto{
				Phone:            phone,
				PassCode:         passCode,
				PhoneCountryCode: phoneCountryCode,
			},
			Options: options,
		},
	})
	var response dto.LoginTokenRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

func (client *Client) SignInByEmailPassCord(email string, passCode string, options dto.SignInOptionsDto) *dto.LoginTokenRespDto {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/signin"),
		Headers: client.getReqHeaders(nil),
		ReqDto: &dto.SigninByCredentialsDto{
			Connection: "PASSCODE",
			PassCodePayload: dto.AuthenticateByPassCodeDto{
				Email:    email,
				PassCode: passCode,
			},
			Options: options,
		},
	})
	var response dto.LoginTokenRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

func (client *Client) SignInByLDAP(sAMAccountName string, passCode string, options dto.SignInOptionsDto) *dto.LoginTokenRespDto {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/signin"),
		Headers: client.getReqHeaders(nil),
		ReqDto: &dto.SigninByCredentialsDto{
			Connection: "LDAP",
			LdapPayload: dto.AuthenticateByLDAPDto{
				SAMAccountName: sAMAccountName,
				Password:       passCode,
			},
			Options: options,
		},
	})
	var response dto.LoginTokenRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

func (client *Client) SignInByAD(sAMAccountName string, passCode string, options dto.SignInOptionsDto) *dto.LoginTokenRespDto {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/signin"),
		Headers: client.getReqHeaders(nil),
		ReqDto: &dto.SigninByCredentialsDto{
			Connection: "AD",
			AdPayload: dto.AuthenticateByADDto{
				SAMAccountName: sAMAccountName,
				Password:       passCode,
			},
			Options: options,
		},
	})
	var response dto.LoginTokenRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

func (client *Client) SignUpByEmailPassCode(email string, passCode string, options dto.SignupOptionsDto) *dto.UserSingleRespDto {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/signup"),
		Headers: client.getReqHeaders(nil),
		ReqDto: &dto.SignupDto{
			Connection: "PASSCODE",
			PassCodePayload: dto.SignUpByPassCodeDto{
				PassCode: passCode,
				Email:    email,
			},
			Options: options,
		},
	})
	var response dto.UserSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

func (client *Client) SignUpByPhonePassCode(phone string, passCode string, phoneCountryCode string, options dto.SignupOptionsDto) *dto.UserSingleRespDto {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/signup"),
		Headers: client.getReqHeaders(nil),
		ReqDto: &dto.SignupDto{
			Connection: "PASSCODE",
			PassCodePayload: dto.SignUpByPassCodeDto{
				PassCode:         passCode,
				Phone:            phone,
				PhoneCountryCode: phoneCountryCode,
			},
			Options: options,
		},
	})
	var response dto.UserSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

func (client *Client) SignUpByEmailPassword(email string, password string, options dto.SignupOptionsDto) *dto.UserSingleRespDto {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/signup"),
		Headers: client.getReqHeaders(nil),
		ReqDto: &dto.SignupDto{
			Connection: "PASSWORD",
			PasswordPayload: dto.SignUpByPasswordDto{
				Email:    email,
				Password: password,
			},
			Options: options,
		},
	})
	var response dto.UserSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

// ==== AUTO GENERATED AUTHENTICATION METHODS BEGIN ====
/*
	* @summary 发起绑定 MFA 认证要素请求
	* @description 当用户未绑定某个 MFA 认证要素时，可以发起绑定 MFA 认证要素请求。不同类型的 MFA 认证要素绑定请求需要发送不同的参数，详细见 profile 参数。发起验证请求之后，Authing 服务器会根据相应的认证要素类型和传递的参数，使用不同的手段要求验证。此接口会返回 enrollmentToken，你需要在请求「绑定 MFA 认证要素」接口时带上此 enrollmentToken，并提供相应的凭证。
	* @returns SendEnrollFactorRequestRespDto 
	*/
func (client *Client) SendEnrollFactorRequest (reqDto *dto.SendEnrollFactorRequestDto) *dto.SendEnrollFactorRequestRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/send-enroll-factor-request"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.SendEnrollFactorRequestRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 绑定 MFA 认证要素
	* @description 绑定 MFA 要素
	* @returns EnrollFactorRespDto 
	*/
func (client *Client) EnrollFactor (reqDto *dto.EnrollFactorDto) *dto.EnrollFactorRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/enroll-factor"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.EnrollFactorRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 解绑 MFA 认证要素
	* @description 当前不支持通过此接口解绑短信、邮箱验证码类型的认证要素。如果需要，请调用「解绑邮箱」和「解绑手机号」接口。
	* @returns ResetFactorRespDto 
	*/
func (client *Client) ResetFactor (reqDto *dto.RestFactorDto) *dto.ResetFactorRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/reset-factor"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.ResetFactorRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 获取绑定的所有 MFA 认证要素
	* @description Authing 目前支持四种类型的 MFA 认证要素：手机短信、邮件验证码、OTP、人脸。如果用户绑定了手机号 / 邮箱之后，默认就具备了手机短信、邮箱验证码的 MFA 认证要素。
	* @returns ListEnrolledFactorsRespDto 
	*/
func (client *Client) ListEnrolledFactors () *dto.ListEnrolledFactorsRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodGet,
		Url:     client.getUrl("/api/v3/list-enrolled-factors"),
		Headers: client.getReqHeaders(nil),
		ReqDto: nil,
	})
	var response dto.ListEnrolledFactorsRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 获取绑定的某个 MFA 认证要素
	* @description 根据 Factor ID 获取用户绑定的某个 MFA Factor 详情。
	* @returns GetFactorRespDto 
	*/
func (client *Client) GetFactor (reqDto *dto.GetFactorDto) *dto.GetFactorRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodGet,
		Url:     client.getUrl("/api/v3/get-factor"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.GetFactorRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 获取可绑定的 MFA 认证要素
	* @description 获取所有应用已经开启、用户暂未绑定的 MFA 认证要素，用户可以从返回的列表中绑定新的 MFA 认证要素。
	* @returns ListFactorsToEnrollRespDto 
	*/
func (client *Client) ListFactorsToEnroll () *dto.ListFactorsToEnrollRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodGet,
		Url:     client.getUrl("/api/v3/list-factors-to-enroll"),
		Headers: client.getReqHeaders(nil),
		ReqDto: nil,
	})
	var response dto.ListFactorsToEnrollRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 绑定外部身份源
	* @description 
 * 
 * 由于绝大多数的外部身份源登录不允许在第三方系统直接输入账号密码进行登录，所以外部身份源的绑定总是需要先跳转到对方的登录页面进行认证。此端点会通过浏览器 `302` 跳转的方式先跳转到第三方的登录页面，
 * 终端用户在第三方系统认证完成之后，浏览器再会跳转到 Authing 服务器，Authing 服务器会将此外部身份源绑定到该用户身上。最终的结果会通过浏览器 Window Post Message 的方式传递给开发者。
 * 你可以在你的应用系统中放置一个按钮，引导用户点击之后，弹出一个 Window Popup，地址为此端点，当用户在第三方身份源认证完成之后，此 Popup 会通过 Window Post Message 的方式传递给父窗口。
 * 
 * 为此我们在 `@authing/browser` SDK 中封装了相关方法，为开发者省去了其中大量的细节：
 * 
 * ```typescript
 * import { Authing } from "@authing/browser"
 * const sdk = new Authing({
 * // 应用的认证地址，例如：https://domain.authing.cn
 * domain: "",
 * 
 * // Authing 应用 ID
 * appId: "you_authing_app_id",
 * 
 * // 登录回调地址，需要在控制台『应用配置 - 登录回调 URL』中指定
 * redirectUri: "your_redirect_uri"
 * });
 * 
 * 
 * // success 表示此次绑定操作是否成功；
 * // errMsg 为如果绑定失败，具体的失败原因，如此身份源已被其他账号绑定等。
 * // identities 为此次绑定操作具体绑定的第三方身份信息
 * const { success, errMsg, identities } = await sdk.bindExtIdpWithPopup({
 * "extIdpConnIdentifier": "my-wechat"
 * })
 * 
 * ```
 * 
 * 绑定外部身份源成功之后，你可以得到用户在此第三方身份源的信息，以绑定飞书账号为例：
 * 
 * ```json
 * [
 * {
 * "identityId": "62f20932xxxxbcc10d966ee5",
 * "extIdpId": "62f209327xxxxcc10d966ee5",
 * "provider": "lark",
 * "type": "open_id",
 * "userIdInIdp": "ou_8bae746eac07cd2564654140d2a9ac61",
 * "originConnIds": ["62f2093244fa5cb19ff21ed3"]
 * },
 * {
 * "identityId": "62f726239xxxxe3285d21c93",
 * "extIdpId": "62f209327xxxxcc10d966ee5",
 * "provider": "lark",
 * "type": "union_id",
 * "userIdInIdp": "on_093ce5023288856aa0abe4099123b18b",
 * "originConnIds": ["62f2093244fa5cb19ff21ed3"]
 * },
 * {
 * "identityId": "62f72623e011cf10c8851e4c",
 * "extIdpId": "62f209327xxxxcc10d966ee5",
 * "provider": "lark",
 * "type": "user_id",
 * "userIdInIdp": "23ded785",
 * "originConnIds": ["62f2093244fa5cb19ff21ed3"]
 * }
 * ]
 * ```
 * 
 * 可以看到，我们获取到了用户在飞书中的身份信息：
 * 
 * - `open_id`: ou_8bae746eac07cd2564654140d2a9ac61
 * - `union_id`: on_093ce5023288856aa0abe4099123b18b
 * - `user_id`: 23ded785
 * 
 * 绑定此外部身份源之后，后续用户就可以使用此身份源进行登录了，见**登录**接口。
 * 
 * 
	* @returns GenerateBindExtIdpLinkRespDto 
	*/
func (client *Client) LinkExtIdp (reqDto *dto.LinkExtidpDto) *dto.GenerateBindExtIdpLinkRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodGet,
		Url:     client.getUrl("/api/v3/link-extidp"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.GenerateBindExtIdpLinkRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 生成绑定外部身份源的链接
	* @description 
 * 
 * 由于绝大多数的外部身份源登录不允许在第三方系统直接输入账号密码进行登录，所以外部身份源的绑定总是需要先跳转到对方的登录页面进行认证。此端点会通过浏览器 `302` 跳转的方式先跳转到第三方的登录页面，
 * 终端用户在第三方系统认证完成之后，浏览器再会跳转到 Authing 服务器，Authing 服务器会将此外部身份源绑定到该用户身上。最终的结果会通过浏览器 Window Post Message 的方式传递给开发者。
 * 你可以在你的应用系统中放置一个按钮，引导用户点击之后，弹出一个 Window Popup，地址为此端点，当用户在第三方身份源认证完成之后，此 Popup 会通过 Window Post Message 的方式传递给父窗口。
 * 
 * 为此我们在 `@authing/browser` SDK 中封装了相关方法，为开发者省去了其中大量的细节：
 * 
 * ```typescript
 * import { Authing } from "@authing/browser"
 * const sdk = new Authing({
 * // 应用的认证地址，例如：https://domain.authing.cn
 * domain: "",
 * 
 * // Authing 应用 ID
 * appId: "you_authing_app_id",
 * 
 * // 登录回调地址，需要在控制台『应用配置 - 登录回调 URL』中指定
 * redirectUri: "your_redirect_uri"
 * });
 * 
 * 
 * // success 表示此次绑定操作是否成功；
 * // errMsg 为如果绑定失败，具体的失败原因，如此身份源已被其他账号绑定等。
 * // identities 为此次绑定操作具体绑定的第三方身份信息
 * const { success, errMsg, identities } = await sdk.bindExtIdpWithPopup({
 * "extIdpConnIdentifier": "my-wechat"
 * })
 * 
 * ```
 * 
 * 绑定外部身份源成功之后，你可以得到用户在此第三方身份源的信息，以绑定飞书账号为例：
 * 
 * ```json
 * [
 * {
 * "identityId": "62f20932xxxxbcc10d966ee5",
 * "extIdpId": "62f209327xxxxcc10d966ee5",
 * "provider": "lark",
 * "type": "open_id",
 * "userIdInIdp": "ou_8bae746eac07cd2564654140d2a9ac61",
 * "originConnIds": ["62f2093244fa5cb19ff21ed3"]
 * },
 * {
 * "identityId": "62f726239xxxxe3285d21c93",
 * "extIdpId": "62f209327xxxxcc10d966ee5",
 * "provider": "lark",
 * "type": "union_id",
 * "userIdInIdp": "on_093ce5023288856aa0abe4099123b18b",
 * "originConnIds": ["62f2093244fa5cb19ff21ed3"]
 * },
 * {
 * "identityId": "62f72623e011cf10c8851e4c",
 * "extIdpId": "62f209327xxxxcc10d966ee5",
 * "provider": "lark",
 * "type": "user_id",
 * "userIdInIdp": "23ded785",
 * "originConnIds": ["62f2093244fa5cb19ff21ed3"]
 * }
 * ]
 * ```
 * 
 * 可以看到，我们获取到了用户在飞书中的身份信息：
 * 
 * - `open_id`: ou_8bae746eac07cd2564654140d2a9ac61
 * - `union_id`: on_093ce5023288856aa0abe4099123b18b
 * - `user_id`: 23ded785
 * 
 * 绑定此外部身份源之后，后续用户就可以使用此身份源进行登录了，见**登录**接口。
 * 
 * 
	* @returns GenerateBindExtIdpLinkRespDto 
	*/
func (client *Client) GenerateLinkExtIdpUrl (reqDto *dto.GenerateLinkExtidpUrlDto) *dto.GenerateBindExtIdpLinkRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodGet,
		Url:     client.getUrl("/api/v3/generate-link-extidp-url"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.GenerateBindExtIdpLinkRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 解绑外部身份源
	* @description 解绑外部身份源，此接口需要传递用户绑定的外部身份源 ID，**注意不是身份源连接 ID**。
	* @returns CommonResponseDto 
	*/
func (client *Client) UnbindExtIdp (reqDto *dto.UnbindExtIdpDto) *dto.CommonResponseDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/unlink-extidp"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.CommonResponseDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 获取绑定的外部身份源
	* @description 
 * 如在**介绍**部分中所描述的，一个外部身份源对应多个外部身份源连接，用户通过某个外部身份源连接绑定了某个外部身份源账号之后，
 * 用户会建立一条与此外部身份源之间的关联关系。此接口用于获取此用户绑定的所有外部身份源。
 * 
 * 取决于外部身份源的具体实现，一个用户在外部身份源中，可能会有多个身份 ID，比如在微信体系中会有 `openid` 和 `unionid`，在非书中有
 * `open_id`、`union_id` 和 `user_id`。在 Authing 中，我们把这样的一条 `open_id` 或者 `unionid_` 叫做一条 `Identity`， 所以用户在一个身份源会有多条 `Identity` 记录。
 * 
 * 以微信为例，如果用户使用微信登录或者绑定了微信账号，他的 `Identity` 信息如下所示：
 * 
 * ```json
 * [
 * {
 * "identityId": "62f20932xxxxbcc10d966ee5",
 * "extIdpId": "62f209327xxxxcc10d966ee5",
 * "provider": "wechat",
 * "type": "openid",
 * "userIdInIdp": "oH_5k5SflrwjGvk7wqpoBKq_cc6M",
 * "originConnIds": ["62f2093244fa5cb19ff21ed3"]
 * },
 * {
 * "identityId": "62f726239xxxxe3285d21c93",
 * "extIdpId": "62f209327xxxxcc10d966ee5",
 * "provider": "wechat",
 * "type": "unionid",
 * "userIdInIdp": "o9Nka5ibU-lUGQaeAHqu0nOZyJg0",
 * "originConnIds": ["62f2093244fa5cb19ff21ed3"]
 * }
 * ]
 * ```
 * 
 * 
 * 可以看到他们的 `extIdpId` 是一样的，这个是你在 Authing 中创建的**身份源 ID**；`provider` 都是 `wechat`；
 * 通过 `type` 可以区分出哪个是 `openid`，哪个是 `unionid`，以及具体的值（`userIdInIdp`）；他们都来自于同一个身份源连接（`originConnIds`）。
 * 
 * 
 * 
	* @returns GetIdentitiesRespDto 
	*/
func (client *Client) GetIdentities () *dto.GetIdentitiesRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodGet,
		Url:     client.getUrl("/api/v3/get-identities"),
		Headers: client.getReqHeaders(nil),
		ReqDto: nil,
	})
	var response dto.GetIdentitiesRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 获取应用开启的外部身份源列表
	* @description 获取应用开启的外部身份源列表，前端可以基于此渲染外部身份源按钮。
	* @returns GetExtIdpsRespDto 
	*/
func (client *Client) GetExtIdps () *dto.GetExtIdpsRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodGet,
		Url:     client.getUrl("/api/v3/get-extidps"),
		Headers: client.getReqHeaders(nil),
		ReqDto: nil,
	})
	var response dto.GetExtIdpsRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 注册
	* @description 
 * 此端点目前支持以下几种基于的注册方式：
 * 
 * 1. 基于密码（PASSWORD）：用户名 + 密码，邮箱 + 密码。
 * 2. 基于一次性临时验证码（PASSCODE）：手机号 + 验证码，邮箱 + 验证码。你需要先调用发送短信或者发送邮件接口获取验证码。
 * 
 * 社会化登录等使用外部身份源“注册”请直接使用**登录**接口，我们会在其第一次登录的时候为其创建一个新账号。
 * 
	* @returns UserSingleRespDto 
	*/
func (client *Client) SignUp (reqDto *dto.SignupDto) *dto.UserSingleRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/signup"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.UserSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 解密微信小程序数据
	* @returns DecryptWechatMiniProgramDataRespDto 
	*/
func (client *Client) DecryptWechatMiniProgramData (reqDto *dto.DecryptWechatMiniProgramDataDto) *dto.DecryptWechatMiniProgramDataRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/decrypt-wechat-miniprogram-data"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.DecryptWechatMiniProgramDataRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 获取小程序的手机号
	* @returns GetWechatMiniProgramPhoneRespDto 
	*/
func (client *Client) GetWechatMiniprogramPhone (reqDto *dto.GetWechatMiniProgramPhoneDto) *dto.GetWechatMiniProgramPhoneRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/get-wechat-miniprogram-phone"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.GetWechatMiniProgramPhoneRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 获取 Authing 服务器缓存的微信小程序、公众号 Access Token
	* @returns GetWechatAccessTokenRespDto 
	*/
func (client *Client) GetWechatMpAccessToken (reqDto *dto.GetWechatAccessTokenDto) *dto.GetWechatAccessTokenRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/get-wechat-access-token"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.GetWechatAccessTokenRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 获取登录日志
	* @description 获取登录日志
	* @returns GetLoginHistoryRespDto 
	*/
func (client *Client) GetLoginHistory (reqDto *dto.GetLoginHistoryDto) *dto.GetLoginHistoryRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodGet,
		Url:     client.getUrl("/api/v3/get-login-history"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.GetLoginHistoryRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 获取登录应用
	* @description 获取登录应用
	* @returns GetLoggedInAppsRespDto 
	*/
func (client *Client) GetLoggedInApps () *dto.GetLoggedInAppsRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodGet,
		Url:     client.getUrl("/api/v3/get-logged-in-apps"),
		Headers: client.getReqHeaders(nil),
		ReqDto: nil,
	})
	var response dto.GetLoggedInAppsRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 获取具备访问权限的应用
	* @description 获取具备访问权限的应用
	* @returns GetAccessibleAppsRespDto 
	*/
func (client *Client) GetAccessibleApps () *dto.GetAccessibleAppsRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodGet,
		Url:     client.getUrl("/api/v3/get-accessible-apps"),
		Headers: client.getReqHeaders(nil),
		ReqDto: nil,
	})
	var response dto.GetAccessibleAppsRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 获取部门列表
	* @description 此接口用于获取用户的部门列表，可根据一定排序规则进行排序。
	* @returns UserDepartmentPaginatedRespDto 
	*/
func (client *Client) GetDepartmentList (reqDto *dto.GetDepartmentListDto) *dto.UserDepartmentPaginatedRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodGet,
		Url:     client.getUrl("/api/v3/get-department-list"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.UserDepartmentPaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 获取被授权的资源列表
	* @description 此接口用于获取用户被授权的资源列表。
	* @returns AuthorizedResourcePaginatedRespDto 
	*/
func (client *Client) GetAuthorizedResources (reqDto *dto.GetAuthorizedResourcesDto) *dto.AuthorizedResourcePaginatedRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodGet,
		Url:     client.getUrl("/api/v3/get-authorized-resources"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.AuthorizedResourcePaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 文件上传
	* @returns UploadRespDto 
	*/
func (client *Client) Upload (reqDto *dto.UploadDto) *dto.UploadRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v2/upload"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.UploadRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 使用用户凭证登录
	* @description 
 * 此端点为基于直接 API 调用形式的登录端点，适用于你需要自建登录页面的场景。**此端点暂时不支持 MFA、信息补全、首次密码重置等流程，如有需要，请使用 OIDC 标准协议认证端点。**
 * 
 * 
 * 注意事项：取决于你在 Authing 创建应用时选择的**应用类型**和应用配置的**换取 token 身份验证方式**，在调用此接口时需要对客户端的身份进行不同形式的验证。
 * 
 * <details>
 * <summary>点击展开详情</summary>
 * 
 * <br>
 * 
 * 你可以在 [Authing 控制台](https://console.authing.cn) 的**应用** - **自建应用** - **应用详情** - **应用配置** - **其他设置** - **授权配置**
 * 中找到**换取 token 身份验证方式** 配置项：
 * 
 * > 单页 Web 应用和客户端应用隐藏，默认为 `none`，不允许修改；后端应用和标准 Web 应用可以修改此配置项。
 * 
 * ![](https://files.authing.co/api-explorer/tokenAuthMethod.jpg)
 * 
 * #### 换取 token 身份验证方式为 none 时
 * 
 * 调用此接口不需要进行额外操作。
 * 
 * #### 换取 token 身份验证方式为 client_secret_post 时
 * 
 * 调用此接口时必须在 body 中传递 `client_id` 和 `client_secret` 参数，作为验证客户端身份的条件。其中 `client_id` 为应用 ID、`client_secret` 为应用密钥。
 * 
 * #### 换取 token 身份验证方式为 client_secret_basic 时
 * 
 * 调用此接口时必须在 HTTP 请求头中携带 `authorization` 请求头，作为验证客户端身份的条件。`authorization` 请求头的格式如下（其中 `client_id` 为应用 ID、`client_secret` 为应用密钥。）：
 * 
 * ```
 * Basic base64(<client_id>:<client_secret>)
 * ```
 * 
 * 结果示例：
 * 
 * ```
 * Basic NjA2M2ZiMmYzY3h4eHg2ZGY1NWYzOWViOjJmZTdjODdhODFmODY3eHh4eDAzMjRkZjEyZGFlZGM3
 * ```
 * 
 * JS 代码示例：
 * 
 * ```js
 * 'Basic ' + Buffer.from(client_id + ':' + client_secret).toString('base64');
 * ```
 * 
 * </details>
 * 
 * 
	* @returns LoginTokenRespDto 成功认证
	*/
func (client *Client) SignInByCredentials (reqDto *dto.SigninByCredentialsDto) *dto.LoginTokenRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/signin"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.LoginTokenRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 使用移动端社会化登录
	* @description 
 * 此端点为移动端社会化登录接口，使用第三方移动社会化登录返回的临时凭证登录，并换取用户的 `id_token` 和 `access_token`。请先阅读相应社会化登录的接入流程。
 * 
 * 
 * 注意事项：取决于你在 Authing 创建应用时选择的**应用类型**和应用配置的**换取 token 身份验证方式**，在调用此接口时需要对客户端的身份进行不同形式的验证。
 * 
 * <details>
 * <summary>点击展开详情</summary>
 * 
 * <br>
 * 
 * 你可以在 [Authing 控制台](https://console.authing.cn) 的**应用** - **自建应用** - **应用详情** - **应用配置** - **其他设置** - **授权配置**
 * 中找到**换取 token 身份验证方式** 配置项：
 * 
 * > 单页 Web 应用和客户端应用隐藏，默认为 `none`，不允许修改；后端应用和标准 Web 应用可以修改此配置项。
 * 
 * ![](https://files.authing.co/api-explorer/tokenAuthMethod.jpg)
 * 
 * #### 换取 token 身份验证方式为 none 时
 * 
 * 调用此接口不需要进行额外操作。
 * 
 * #### 换取 token 身份验证方式为 client_secret_post 时
 * 
 * 调用此接口时必须在 body 中传递 `client_id` 和 `client_secret` 参数，作为验证客户端身份的条件。其中 `client_id` 为应用 ID、`client_secret` 为应用密钥。
 * 
 * #### 换取 token 身份验证方式为 client_secret_basic 时
 * 
 * 调用此接口时必须在 HTTP 请求头中携带 `authorization` 请求头，作为验证客户端身份的条件。`authorization` 请求头的格式如下（其中 `client_id` 为应用 ID、`client_secret` 为应用密钥。）：
 * 
 * ```
 * Basic base64(<client_id>:<client_secret>)
 * ```
 * 
 * 结果示例：
 * 
 * ```
 * Basic NjA2M2ZiMmYzY3h4eHg2ZGY1NWYzOWViOjJmZTdjODdhODFmODY3eHh4eDAzMjRkZjEyZGFlZGM3
 * ```
 * 
 * JS 代码示例：
 * 
 * ```js
 * 'Basic ' + Buffer.from(client_id + ':' + client_secret).toString('base64');
 * ```
 * 
 * </details>
 * 
 * 
	* @returns LoginTokenRespDto 
	*/
func (client *Client) SignInByMobile (reqDto *dto.SigninByMobileDto) *dto.LoginTokenRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/signin-by-mobile"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.LoginTokenRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 获取支付宝 AuthInfo
	* @description 此接口用于获取发起支付宝认证需要的[初始化参数 AuthInfo](https://opendocs.alipay.com/open/218/105325)。
	* @returns GetAlipayAuthInfoRespDto 
	*/
func (client *Client) GetAlipayAuthInfo (reqDto *dto.GetAlipayAuthinfoDto) *dto.GetAlipayAuthInfoRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodGet,
		Url:     client.getUrl("/api/v3/get-alipay-authinfo"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.GetAlipayAuthInfoRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 生成用于登录的二维码
	* @description 生成用于登录的二维码，目前支持生成微信公众号扫码登录、小程序扫码登录、自建移动 APP 扫码登录的二维码。
	* @returns GeneQRCodeRespDto 
	*/
func (client *Client) GeneQrCode (reqDto *dto.GenerateQrcodeDto) *dto.GeneQRCodeRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/gene-qrcode"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.GeneQRCodeRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 查询二维码状态
	* @description 按照用户扫码顺序，共分为未扫码、已扫码等待用户确认、用户同意/取消授权、二维码过期以及未知错误六种状态，前端应该通过不同的状态给到用户不同的反馈。你可以通过下面这篇文章了解扫码登录详细的流程：https://docs.authing.cn/v2/concepts/how-qrcode-works.html.
	* @returns CheckQRCodeStatusRespDto 
	*/
func (client *Client) CheckQrCodeStatus (reqDto *dto.CheckQrcodeStatusDto) *dto.CheckQRCodeStatusRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodGet,
		Url:     client.getUrl("/api/v3/check-qrcode-status"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.CheckQRCodeStatusRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 使用二维码 ticket 换取 TokenSet
	* @description 
 * 此端点为使用二维码的 ticket 换取用户的 `access_token` 和 `id_token`。
 * 
 * 
 * 注意事项：取决于你在 Authing 创建应用时选择的**应用类型**和应用配置的**换取 token 身份验证方式**，在调用此接口时需要对客户端的身份进行不同形式的验证。
 * 
 * <details>
 * <summary>点击展开详情</summary>
 * 
 * <br>
 * 
 * 你可以在 [Authing 控制台](https://console.authing.cn) 的**应用** - **自建应用** - **应用详情** - **应用配置** - **其他设置** - **授权配置**
 * 中找到**换取 token 身份验证方式** 配置项：
 * 
 * > 单页 Web 应用和客户端应用隐藏，默认为 `none`，不允许修改；后端应用和标准 Web 应用可以修改此配置项。
 * 
 * ![](https://files.authing.co/api-explorer/tokenAuthMethod.jpg)
 * 
 * #### 换取 token 身份验证方式为 none 时
 * 
 * 调用此接口不需要进行额外操作。
 * 
 * #### 换取 token 身份验证方式为 client_secret_post 时
 * 
 * 调用此接口时必须在 body 中传递 `client_id` 和 `client_secret` 参数，作为验证客户端身份的条件。其中 `client_id` 为应用 ID、`client_secret` 为应用密钥。
 * 
 * #### 换取 token 身份验证方式为 client_secret_basic 时
 * 
 * 调用此接口时必须在 HTTP 请求头中携带 `authorization` 请求头，作为验证客户端身份的条件。`authorization` 请求头的格式如下（其中 `client_id` 为应用 ID、`client_secret` 为应用密钥。）：
 * 
 * ```
 * Basic base64(<client_id>:<client_secret>)
 * ```
 * 
 * 结果示例：
 * 
 * ```
 * Basic NjA2M2ZiMmYzY3h4eHg2ZGY1NWYzOWViOjJmZTdjODdhODFmODY3eHh4eDAzMjRkZjEyZGFlZGM3
 * ```
 * 
 * JS 代码示例：
 * 
 * ```js
 * 'Basic ' + Buffer.from(client_id + ':' + client_secret).toString('base64');
 * ```
 * 
 * </details>
 * 
 * 
	* @returns LoginTokenRespDto 
	*/
func (client *Client) ExchangeTokenSetWithQrCodeTicket (reqDto *dto.ExchangeTokenSetWithQRcodeTicketDto) *dto.LoginTokenRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/exchange-tokenset-with-qrcode-ticket"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.LoginTokenRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 自建 APP 扫码登录：APP 端修改二维码状态
	* @description 此端点用于在自建 APP 扫码登录中修改二维码状态，对应着在浏览器渲染出二维码之后，终端用户扫码、确认授权、取消授权的过程。**此接口要求具备用户的登录态**。
	* @returns CommonResponseDto 
	*/
func (client *Client) ChangeQrCodeStatus (reqDto *dto.ChangeQRCodeStatusDto) *dto.CommonResponseDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/change-qrcode-status"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.CommonResponseDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 发送短信
	* @description 发送短信时必须指定短信 Channel，每个手机号同一 Channel 在一分钟内只能发送一次。
	* @returns SendSMSRespDto 
	*/
func (client *Client) SendSms (reqDto *dto.SendSMSDto) *dto.SendSMSRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/send-sms"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.SendSMSRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 发送邮件
	* @description 发送邮件时必须指定邮件 Channel，每个邮箱同一 Channel 在一分钟内只能发送一次。
	* @returns SendEmailRespDto 
	*/
func (client *Client) SendEmail (reqDto *dto.SendEmailDto) *dto.SendEmailRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/send-email"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.SendEmailRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 获取用户资料
	* @description 此端点用户获取用户资料，需要在请求头中带上用户的 `access_token`，Authing 服务器会根据用户 `access_token` 中的 `scope` 返回对应的字段。
	* @returns UserSingleRespDto 
	*/
func (client *Client) GetProfile (reqDto *dto.GetProfileDto) *dto.UserSingleRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodGet,
		Url:     client.getUrl("/api/v3/get-profile"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.UserSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 修改用户资料
	* @description 此接口用于修改用户的用户资料，包含用户的自定义数据。如果需要**修改邮箱**、**修改手机号**、**修改密码**，请使用对应的单独接口。
	* @returns UserSingleRespDto 
	*/
func (client *Client) UpdateProfile (reqDto *dto.UpdateUserProfileDto) *dto.UserSingleRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/update-profile"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.UserSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 绑定邮箱
	* @description 如果用户还**没有绑定邮箱**，此接口可用于用户**自主**绑定邮箱。如果用户已经绑定邮箱想要修改邮箱，请使用**修改邮箱**接口。你需要先调用**发送邮件**接口发送邮箱验证码。
	* @returns CommonResponseDto 
	*/
func (client *Client) BindEmail (reqDto *dto.BindEmailDto) *dto.CommonResponseDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/bind-email"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.CommonResponseDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 解绑邮箱
	* @description 用户解绑邮箱，如果用户没有绑定其他登录方式（手机号、社会化登录账号），将无法解绑邮箱，会提示错误。
	* @returns CommonResponseDto 
	*/
func (client *Client) UnbindEmail (reqDto *dto.UnbindEmailDto) *dto.CommonResponseDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/unbind-email"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.CommonResponseDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 绑定手机号
	* @description 如果用户还**没有绑定手机号**，此接口可用于用户**自主**绑定手机号。如果用户已经绑定手机号想要修改手机号，请使用**修改手机号**接口。你需要先调用**发送短信**接口发送短信验证码。
	* @returns CommonResponseDto 
	*/
func (client *Client) BindPhone (reqDto *dto.BindPhoneDto) *dto.CommonResponseDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/bind-phone"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.CommonResponseDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 解绑手机号
	* @description 用户解绑手机号，如果用户没有绑定其他登录方式（邮箱、社会化登录账号），将无法解绑手机号，会提示错误。
	* @returns CommonResponseDto 
	*/
func (client *Client) UnbindPhone (reqDto *dto.UnbindPhoneDto) *dto.CommonResponseDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/unbind-phone"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.CommonResponseDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 获取密码强度和账号安全等级评分
	* @description 获取用户的密码强度和账号安全等级评分，需要在请求头中带上用户的 `access_token`。
	* @returns GetSecurityInfoRespDto 
	*/
func (client *Client) GetSecurityLevel () *dto.GetSecurityInfoRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodGet,
		Url:     client.getUrl("/api/v3/get-security-info"),
		Headers: client.getReqHeaders(nil),
		ReqDto: nil,
	})
	var response dto.GetSecurityInfoRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 修改密码
	* @description 此端点用于用户自主修改密码，如果用户之前已经设置密码，需要提供用户的原始密码作为凭证。如果用户忘记了当前密码，请使用**忘记密码**接口。
	* @returns CommonResponseDto 
	*/
func (client *Client) UpdatePassword (reqDto *dto.UpdatePasswordDto) *dto.CommonResponseDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/update-password"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.CommonResponseDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 发起修改邮箱的验证请求
	* @description 终端用户自主修改邮箱时，需要提供相应的验证手段。此接口用于验证用户的修改邮箱请求是否合法。当前支持通过**邮箱验证码**的方式进行验证，你需要先调用发送邮件接口发送对应的邮件验证码。
	* @returns VerifyUpdateEmailRequestRespDto 
	*/
func (client *Client) VerifyUpdateEmailRequest (reqDto *dto.VerifyUpdateEmailRequestDto) *dto.VerifyUpdateEmailRequestRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/verify-update-email-request"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.VerifyUpdateEmailRequestRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 修改邮箱
	* @description 终端用户自主修改邮箱，需要提供相应的验证手段，见[发起修改邮箱的验证请求](#tag/用户资料/修改邮箱/operation/ProfileV3Controller_updateEmailVerification)。
 * 此参数需要提供一次性临时凭证 `updateEmailToken`，此数据需要从**发起修改邮箱的验证请求**接口获取。
	* @returns CommonResponseDto 
	*/
func (client *Client) UpdateEmail (reqDto *dto.UpdateEmailDto) *dto.CommonResponseDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/update-email"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.CommonResponseDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 发起修改手机号的验证请求
	* @description 终端用户自主修改手机号时，需要提供相应的验证手段。此接口用于验证用户的修改手机号请求是否合法。当前支持通过**短信验证码**的方式进行验证，你需要先调用发送短信接口发送对应的短信验证码。
	* @returns VerifyUpdatePhoneRequestRespDto 
	*/
func (client *Client) VerifyUpdatePhoneRequest (reqDto *dto.VerifyUpdatePhoneRequestDto) *dto.VerifyUpdatePhoneRequestRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/verify-update-phone-request"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.VerifyUpdatePhoneRequestRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 修改手机号
	* @description 终端用户自主修改手机号，需要提供相应的验证手段，见[发起修改手机号的验证请求](#tag/用户资料/修改邮箱/operation/ProfileV3Controller_updatePhoneVerification)。
 * 此参数需要提供一次性临时凭证 `updatePhoneToken`，此数据需要从**发起修改手机号的验证请求**接口获取。
	* @returns CommonResponseDto 
	*/
func (client *Client) UpdatePhone (reqDto *dto.UpdatePhoneDto) *dto.CommonResponseDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/update-phone"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.CommonResponseDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 发起忘记密码请求
	* @description 当用户忘记密码时，可以通过此端点找回密码。用户需要使用相关验证手段进行验证，目前支持**邮箱验证码**和**手机号验证码**两种验证手段。
	* @returns PasswordResetVerifyResp 
	*/
func (client *Client) VerifyResetPasswordRequest (reqDto *dto.VerifyResetPasswordRequestDto) *dto.PasswordResetVerifyResp  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/verify-reset-password-request"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.PasswordResetVerifyResp
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 忘记密码
	* @description 此端点用于用户忘记密码之后，通过**手机号验证码**或者**邮箱验证码**的方式重置密码。此接口需要提供用于重置密码的临时凭证 `passwordResetToken`，此参数需要通过**发起忘记密码请求**接口获取。
	* @returns IsSuccessRespDto 
	*/
func (client *Client) ResetPassword (reqDto *dto.ResetPasswordDto) *dto.IsSuccessRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/reset-password"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 发起注销账号请求
	* @description 当用户希望注销账号时，需提供相应凭证，当前支持**使用邮箱验证码**、使用**手机验证码**、**使用密码**三种验证方式。
	* @returns VerifyDeleteAccountRequestRespDto 
	*/
func (client *Client) VeirfyDeleteAccountRequest (reqDto *dto.VerifyDeleteAccountRequestDto) *dto.VerifyDeleteAccountRequestRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/verify-delete-account-request"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.VerifyDeleteAccountRequestRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 注销账户
	* @description 此端点用于用户自主注销账号，需要提供用于注销账号的临时凭证 deleteAccountToken，此参数需要通过**发起注销账号请求**接口获取。
	* @returns IsSuccessRespDto 
	*/
func (client *Client) DeleteAccount (reqDto *dto.DeleteAccounDto) *dto.IsSuccessRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/delete-account"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 获取服务器公开信息
	* @description 可端点可获取服务器的公开信息，如 RSA256 公钥、SM2 公钥、Authing 服务版本号等。
	* @returns SystemInfoResp 
	*/
func (client *Client) GetSystemInfo () *dto.SystemInfoResp  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodGet,
		Url:     client.getUrl("/api/v3/system"),
		Headers: client.getReqHeaders(nil),
		ReqDto: nil,
	})
	var response dto.SystemInfoResp
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 获取国家列表
	* @description 动态获取国家列表，可以用于前端登录页面国家选择和国际短信输入框选择，以减少前端静态资源体积。
	* @returns GetCountryListRespDto 
	*/
func (client *Client) GetCountryList () *dto.GetCountryListRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodGet,
		Url:     client.getUrl("/api/v3/get-country-list"),
		Headers: client.getReqHeaders(nil),
		ReqDto: nil,
	})
	var response dto.GetCountryListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
/*
	* @summary 预检验验证码是否正确
	* @description 预检测验证码是否有效，此检验不会使得验证码失效。
	* @returns PreCheckCodeRespDto 
	*/
func (client *Client) PreCheckCode (reqDto *dto.PreCheckCodeDto) *dto.PreCheckCodeRespDto  {
	res, err := util.SendRequest(&util.RequestOption{
		Method:  fasthttp.MethodPost,
		Url:     client.getUrl("/api/v3/pre-check-code"),
		Headers: client.getReqHeaders(nil),
		ReqDto: reqDto,
	})
	var response dto.PreCheckCodeRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(res.Body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

// ==== AUTO GENERATED AUTHENTICATION METHODS END ====
