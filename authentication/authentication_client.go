package authentication

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Authing/authing-golang-sdk/constant"
	"github.com/Authing/authing-golang-sdk/dto"
	"github.com/Authing/authing-golang-sdk/util"
	"strings"

	keyfunc "github.com/MicahParks/compatibility-keyfunc"
	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
)

const RandStringLen = 16
const ALG_HS256 = "HS256"
const JWK_PATH = "/oidc/.well-known/jwks.json"

var commonHeaders = map[string]string{
	"x-authing-request-from": constant.SdkName,
	"x-authing-sdk-version":  constant.SdkVersion,
}

type AuthenticationClient struct {
	options *AuthenticationClientOptions
	jwks    *keyfunc.JWKS
}

func NewAuthenticationClient(options *AuthenticationClientOptions) (*AuthenticationClient, error) {
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
	if options.Scope == "" {
		options.Scope = constant.DefaultScope
	}
	if options.Protocol == "" {
		options.Protocol = OIDC
	}
	if options.Timeout == 0 {
		options.Timeout = 10000
	}
	if options.TokenEndPointAuthMethod == "" {
		options.TokenEndPointAuthMethod = ClientSecretPost
	}

	client := &AuthenticationClient{
		options: options,
	}

	return client, nil
}

func (client *AuthenticationClient) SetAccessToken(accessToken string) {
	client.options.AccessToken = accessToken
}

func (client *AuthenticationClient) getUrl(path string) string {
	return client.options.AppHost + path
}

func (client *AuthenticationClient) BuildAuthorizeUrlByOidc(params *OIDCAuthURLParams) (AuthUrlResult, error) {
	if params == nil {
		params = &OIDCAuthURLParams{}
	}

	// scope
	scope := params.Scope
	if scope == "" {
		scope = client.options.Scope
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
		redirectUri = client.options.RedirectUri
	}
	responseMode := params.ResponseMode
	if responseMode == "" {
		responseMode = "query"
	}

	dataMap := map[string]interface{}{
		"redirect_uri":  redirectUri,
		"client_id":     client.options.AppId,
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

func (client *AuthenticationClient) BuildAuthorizeUrlByOauth(params *OAuth2AuthURLParams) (string, error) {
	dataMap := map[string]string{
		"client_id":     util.GetValueOrDefault(client.options.AppId),
		"scope":         util.GetValueOrDefault(params.Scope, "openid profile email phone address"),
		"state":         util.GetValueOrDefault(params.State, util.RandomString(12)),
		"response_type": util.GetValueOrDefault(params.ResponseType, "code"),
		"redirect_uri":  util.GetValueOrDefault(params.RedirectUri, client.options.RedirectUri),
	}
	return client.getUrl("/oauth/auth?") + util.GetQueryString(dataMap), nil
}

func (client *AuthenticationClient) BuildAuthorizeUrlBySaml() string {
	return fmt.Sprintf("%s/api/v2/saml-idp/%s", client.options.AppHost, client.options.AppId)
}

func (client *AuthenticationClient) BuildAuthorizeUrlByCas(service *string) string {
	if service != nil {
		return fmt.Sprintf("%s/cas-idp/%s?service=%s", client.options.AppHost, client.options.AppId, *service)
	} else {
		return fmt.Sprintf("%s/cas-idp/%s?service", client.options.AppHost, client.options.AppId)
	}
}

/*
*
检验 CAS 1.0 Ticket 合法性
*/
func (client *AuthenticationClient) ValidateTicketV1(ticket, service string) (*struct {
	Valid    bool   `json:"code"`
	Message  string `json:"message"`
	Username string `json:"username"`
}, error) {
	url := fmt.Sprintf("%s/cas-idp/%s/validate?service=%s&ticket=%s", client.options.AppHost, client.options.AppId, service, ticket)
	res, err := client.SendProtocolHttpRequest(&ProtocolRequestOption{
		Url:     url,
		Method:  fasthttp.MethodGet,
		Headers: client.getReqHeaders(nil),
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
func (client *AuthenticationClient) ValidateTicketV2(ticket, service string, format string) (*struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}, error) {
	url := fmt.Sprintf("%s/cas-idp/%s/serviceValidate", client.options.AppHost, client.options.AppId)
	res, err := client.SendProtocolHttpRequest(&ProtocolRequestOption{
		Url:     url,
		Method:  fasthttp.MethodGet,
		Headers: client.getReqHeaders(nil),
		ReqDto: map[string]string{
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
func (client *AuthenticationClient) GetAccessTokenByCode(code string) (OIDCTokenResponse, error) {
	url := client.options.AppHost + "/oidc/token"
	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	body := map[string]string{
		"client_id":     client.options.AppId,
		"client_secret": client.options.AppSecret,
		"grant_type":    "authorization_code",
		"code":          code,
		"redirect_uri":  client.options.RedirectUri,
	}

	switch client.options.TokenEndPointAuthMethod {
	case ClientSecretPost:
		body["client_id"] = client.options.AppId
		body["client_secret"] = client.options.AppSecret
	case ClientSecretBasic:
		base64String := "Basic " + base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", client.options.AppId, client.options.AppSecret)))
		header["Authorization"] = base64String
	default:
		body["client_id"] = client.options.AppId
	}
	resp, err := client.SendProtocolHttpRequest(&ProtocolRequestOption{
		Url:     url,
		Method:  fasthttp.MethodPost,
		Headers: client.getReqHeaders(header),
		ReqDto:  body,
	})
	var tokenResponse OIDCTokenResponse
	err = json.Unmarshal(resp.Body, &tokenResponse)
	return tokenResponse, err
}

// GetAccessTokenByClientCredentials
// AuthenticationClient Credentials 模式获取 Access Token
func (client *AuthenticationClient) GetAccessTokenByClientCredentials(req GetAccessTokenByClientCredentialsRequest) (string, error) {
	if req.Scope == constant.StringEmpty {
		return constant.StringEmpty, errors.New("请传入 scope 参数，请看文档：https://docs.authing.cn/v2/guides/authorization/m2m-authz.html")
	}
	if req.ClientCredentialInput == nil {
		return constant.StringEmpty, errors.New("请在调用本方法时传入 ClientCredentialInput 参数，请看文档：https://docs.authing.cn/v2/guides/authorization/m2m-authz.html")
	}

	url := client.options.AppHost + "/oidc/token"

	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	body := map[string]string{
		"client_id":     req.ClientCredentialInput.AccessKey,
		"client_secret": req.ClientCredentialInput.SecretKey,
		"grant_type":    "client_credentials",
		"scope":         req.Scope,
	}

	resp, err := client.SendProtocolHttpRequest(&ProtocolRequestOption{
		Url:     url,
		Method:  fasthttp.MethodPost,
		Headers: client.getReqHeaders(header),
		ReqDto:  body,
	})
	return string(resp.Body), err
}

// GetNewAccessTokenByRefreshToken
//
//	使用 Refresh token 获取新的 Access token
func (client *AuthenticationClient) GetNewAccessTokenByRefreshToken(refreshToken string) (string, error) {
	if client.options.Protocol != OIDC && client.options.Protocol != OAUTH {
		return constant.StringEmpty, errors.New("初始化 AuthenticationClient 时传入的 protocol 参数必须为 ProtocolEnum.OAUTH 或 ProtocolEnum.OIDC，请检查参数")
	}
	if client.options.AppSecret == "" && client.options.TokenEndPointAuthMethod != constant.None {
		return constant.StringEmpty, errors.New("请在初始化 AuthenticationClient 时传入 Secret")
	}

	url := client.options.AppHost + fmt.Sprintf("/%s/token", client.options.Protocol)

	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	body := map[string]string{
		"client_id":     client.options.AppId,
		"client_secret": client.options.AppSecret,
		"grant_type":    "refresh_token",
		"refresh_token": refreshToken,
	}

	switch client.options.TokenEndPointAuthMethod {
	case ClientSecretPost:
		body["client_id"] = client.options.AppId
		body["client_secret"] = client.options.AppSecret
	case ClientSecretBasic:
		base64String := "Basic " + base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", client.options.AppId, client.options.AppSecret)))
		header["Authorization"] = base64String
	default:
		body["client_id"] = client.options.AppId
	}
	resp, err := client.SendProtocolHttpRequest(&ProtocolRequestOption{
		Url:     url,
		Method:  fasthttp.MethodPost,
		Headers: client.getReqHeaders(header),
		ReqDto:  body,
	})
	return string(resp.Body), err
}

func (client *AuthenticationClient) IntrospectToken(token string) (string, error) {
	url := client.options.AppHost + fmt.Sprintf("/%s/token/introspection", client.options.Protocol)
	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	body := map[string]string{
		"token": token,
	}
	switch client.options.TokenEndPointAuthMethod {
	case ClientSecretPost:
		body["client_id"] = client.options.AppId
		body["client_secret"] = client.options.AppSecret
	case ClientSecretBasic:
		base64String := "Basic " + base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", client.options.AppId, client.options.AppSecret)))
		header["Authorization"] = base64String
	default:
		body["client_id"] = client.options.AppId
	}
	resp, err := client.SendProtocolHttpRequest(&ProtocolRequestOption{
		Url:     url,
		Method:  fasthttp.MethodPost,
		Headers: client.getReqHeaders(header),
		ReqDto:  body,
	})
	return string(resp.Body), err
}

// RevokeToken
// 撤回 Access token 或 Refresh token
func (client *AuthenticationClient) RevokeToken(token string) (string, error) {
	if client.options.Protocol != OIDC && client.options.Protocol != OAUTH {
		return constant.StringEmpty, errors.New("初始化 AuthenticationClient 时传入的 protocol 参数必须为 ProtocolEnum.OAUTH 或 ProtocolEnum.OIDC，请检查参数")
	}
	if client.options.AppSecret == "" && client.options.TokenEndPointAuthMethod != constant.None {
		return constant.StringEmpty, errors.New("请在初始化 AuthenticationClient 时传入 Secret")
	}

	url := client.options.AppHost + fmt.Sprintf("/%s/token/revocation", client.options.Protocol)

	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	body := map[string]string{
		"client_id": client.options.AppId,
		"token":     token,
	}

	switch client.options.TokenEndPointAuthMethod {
	case ClientSecretPost:
		body["client_id"] = client.options.AppId
		body["client_secret"] = client.options.AppSecret
	case ClientSecretBasic:
		base64String := "Basic " + base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", client.options.AppId, client.options.AppSecret)))
		header["Authorization"] = base64String
	default:
		body["client_id"] = client.options.AppId
	}
	resp, err := client.SendProtocolHttpRequest(&ProtocolRequestOption{
		Url:     url,
		Method:  fasthttp.MethodPost,
		Headers: client.getReqHeaders(header),
		ReqDto:  body,
	})
	return string(resp.Body), err
}

// 拼接登出 URL
func (client *AuthenticationClient) BuildLogoutUrl(params *BuildLogoutURLParams) (string, error) {
	var url string
	if client.options.Protocol == CAS {
		if params.PostLogoutRedirectUri != "" {
			url = fmt.Sprintf("%s/cas-idp/logout?url=%s", client.options.AppHost, params.PostLogoutRedirectUri)
		} else {
			url = fmt.Sprintf("%s/cas-idp/logout", client.options.AppHost)
		}
	} else if client.options.Protocol == OIDC {
		return client.buildLogoutUrlOIDC(params)
	} else {
		if params.PostLogoutRedirectUri != "" {
			url = fmt.Sprintf("%s/login/profile/logout?redirect_uri=%s", client.options.AppHost, params.PostLogoutRedirectUri)
		} else {
			url = fmt.Sprintf("%s/login/profile/logout", client.options.AppHost)
		}
	}
	return url, nil
}

func (client *AuthenticationClient) getJWKS() (*keyfunc.JWKS, error) {

	if client.jwks != nil {
		return client.jwks, nil
	}

	jwksURL := client.getUrl(JWK_PATH)
	jwks, err := keyfunc.Get(jwksURL, keyfunc.Options{})
	if err != nil {
		return nil, fmt.Errorf("获取 jwk 密钥失败: %w", err)
	}
	client.jwks = jwks
	return jwks, err
}

func (client *AuthenticationClient) getKeyCommon(token *jwt.Token) (interface{}, error) {
	alg, ok := token.Header["alg"].(string)
	if !ok {
		return nil, fmt.Errorf("算法字段非法 %v", token.Header["alg"])
	}
	if alg == ALG_HS256 {
		return []byte(client.options.AppSecret), nil
	}
	jwks, err := client.getJWKS()
	if err != nil {
		return nil, fmt.Errorf("获取 JWKS 失败 %v", err)
	}
	return jwks.KeyfuncLegacy(token)
}

func (client *AuthenticationClient) getKey4IdToken(token *jwt.Token) (interface{}, error) {
	claims := token.Claims.(*IDTokenClaims)
	claims.IssuedAt = 0

	return client.getKeyCommon(token)
}

func (client *AuthenticationClient) getKey4AccessToken(token *jwt.Token) (interface{}, error) {
	claims := token.Claims.(*AccessTokenClaims)
	claims.IssuedAt = 0

	return client.getKeyCommon(token)
}

func (client *AuthenticationClient) ParseIDToken(tokenStr string) (*IDTokenClaims, error) {
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

func (client *AuthenticationClient) IntrospectAccessTokenOffline(tokenStr string) (*AccessTokenClaims, error) {
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

func (client *AuthenticationClient) getReqHeaders(customHeaders map[string]string) map[string]string {
	newHeaders := make(map[string]string)
	for key, value := range commonHeaders {
		newHeaders[key] = value
	}
	for key, value := range customHeaders {
		newHeaders[key] = value
	}

	return newHeaders
}
func (client *AuthenticationClient) GetUserInfo(accessToken string) (*UserInfo, error) {
	res, err := client.SendProtocolHttpRequest(&ProtocolRequestOption{
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

func (client *AuthenticationClient) buildLogoutUrlOIDC(params *BuildLogoutURLParams) (string, error) {
	if params == nil {
		params = &BuildLogoutURLParams{}
	}
	idToken := params.IDTokenHint
	redirectUri := params.PostLogoutRedirectUri
	if redirectUri == "" {
		redirectUri = client.options.LogoutRedirectUri
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

func (client *AuthenticationClient) SignInByUsernamePassword(username string, password string, options dto.SignInOptionsDto) *dto.LoginTokenRespDto {
	body, err := client.SendHttpRequest(
		"/api/v3/signin",
		fasthttp.MethodPost,
		&dto.SigninByCredentialsDto{
			Connection: "PASSWORD",
			PasswordPayload: dto.SignInByPasswordPayloadDto{
				Password: password,
				Username: username,
			},
			Options:      options,
			ClientId:     client.options.AppId,
			ClientSecret: client.options.AppSecret,
		},
	)
	var response dto.LoginTokenRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

func (client *AuthenticationClient) SignInByEmailPassword(email string, password string, options dto.SignInOptionsDto) *dto.LoginTokenRespDto {
	body, err := client.SendHttpRequest(
		"/api/v3/signin",
		fasthttp.MethodPost,
		&dto.SigninByCredentialsDto{
			Connection: "PASSWORD",
			PasswordPayload: dto.SignInByPasswordPayloadDto{
				Password: password,
				Email:    email,
			},
			Options:      options,
			ClientId:     client.options.AppId,
			ClientSecret: client.options.AppSecret,
		},
	)
	var response dto.LoginTokenRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

func (client *AuthenticationClient) SignInByPhonePassword(phone string, password string, options dto.SignInOptionsDto) *dto.LoginTokenRespDto {
	body, err := client.SendHttpRequest(
		"/api/v3/signin",
		fasthttp.MethodPost,
		&dto.SigninByCredentialsDto{
			Connection: "PASSWORD",
			PasswordPayload: dto.SignInByPasswordPayloadDto{
				Password: password,
				Phone:    phone,
			},
			Options:      options,
			ClientId:     client.options.AppId,
			ClientSecret: client.options.AppSecret,
		},
	)
	var response dto.LoginTokenRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

func (client *AuthenticationClient) SignInByAccountPassword(account string, password string, options dto.SignInOptionsDto) *dto.LoginTokenRespDto {
	body, err := client.SendHttpRequest(
		"/api/v3/signin",
		fasthttp.MethodPost,
		&dto.SigninByCredentialsDto{
			Connection: "PASSWORD",
			PasswordPayload: dto.SignInByPasswordPayloadDto{
				Password: password,
				Account:  account,
			},
			Options:      options,
			ClientId:     client.options.AppId,
			ClientSecret: client.options.AppSecret,
		},
	)
	var response dto.LoginTokenRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

func (client *AuthenticationClient) SignInByPhonePassCord(phone string, passCode string, phoneCountryCode string, options dto.SignInOptionsDto) *dto.LoginTokenRespDto {
	body, err := client.SendHttpRequest(
		"/api/v3/signin",
		fasthttp.MethodPost,
		&dto.SigninByCredentialsDto{
			Connection: "PASSCODE",
			PassCodePayload: dto.SignInByPassCodePayloadDto{
				Phone:            phone,
				PassCode:         passCode,
				PhoneCountryCode: phoneCountryCode,
			},
			Options:      options,
			ClientId:     client.options.AppId,
			ClientSecret: client.options.AppSecret,
		},
	)
	var response dto.LoginTokenRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

func (client *AuthenticationClient) SignInByEmailPassCord(email string, passCode string, options dto.SignInOptionsDto) *dto.LoginTokenRespDto {
	body, err := client.SendHttpRequest(
		"/api/v3/signin",
		fasthttp.MethodPost,
		&dto.SigninByCredentialsDto{
			Connection: "PASSCODE",
			PassCodePayload: dto.SignInByPassCodePayloadDto{
				Email:    email,
				PassCode: passCode,
			},
			Options:      options,
			ClientId:     client.options.AppId,
			ClientSecret: client.options.AppSecret,
		},
	)
	var response dto.LoginTokenRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

func (client *AuthenticationClient) SignInByLDAP(sAMAccountName string, passCode string, options dto.SignInOptionsDto) *dto.LoginTokenRespDto {
	body, err := client.SendHttpRequest(
		"/api/v3/signin",
		fasthttp.MethodPost,
		&dto.SigninByCredentialsDto{
			Connection: "LDAP",
			LdapPayload: dto.SignInByLdapPayloadDto{
				SAMAccountName: sAMAccountName,
				Password:       passCode,
			},
			Options:      options,
			ClientId:     client.options.AppId,
			ClientSecret: client.options.AppSecret,
		},
	)
	var response dto.LoginTokenRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

func (client *AuthenticationClient) SignInByAD(sAMAccountName string, passCode string, options dto.SignInOptionsDto) *dto.LoginTokenRespDto {
	body, err := client.SendHttpRequest(
		"/api/v3/signin",
		fasthttp.MethodPost,
		&dto.SigninByCredentialsDto{
			Connection: "AD",
			AdPayload: dto.SignInByAdPayloadDto{
				SAMAccountName: sAMAccountName,
				Password:       passCode,
			},
			Options:      options,
			ClientId:     client.options.AppId,
			ClientSecret: client.options.AppSecret,
		},
	)
	var response dto.LoginTokenRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

func (client *AuthenticationClient) SignUpByEmailPassCode(email string, passCode string, options dto.SignUpOptionsDto) *dto.UserSingleRespDto {
	body, err := client.SendHttpRequest(
		"/api/v3/signup",
		fasthttp.MethodPost,
		&dto.SignUpDto{
			Connection: "PASSCODE",
			PassCodePayload: dto.SignUpByPassCodeDto{
				PassCode: passCode,
				Email:    email,
			},
			Options: options,
		},
	)
	var response dto.UserSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

func (client *AuthenticationClient) SignUpByPhonePassCode(phone string, passCode string, phoneCountryCode string, options dto.SignUpOptionsDto) *dto.UserSingleRespDto {
	body, err := client.SendHttpRequest(
		"/api/v3/signup",
		fasthttp.MethodPost,
		&dto.SignUpDto{
			Connection: "PASSCODE",
			PassCodePayload: dto.SignUpByPassCodeDto{
				PassCode:         passCode,
				Phone:            phone,
				PhoneCountryCode: phoneCountryCode,
			},
			Options: options,
		},
	)
	var response dto.UserSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

func (client *AuthenticationClient) SignUpByEmailPassword(email string, password string, options dto.SignUpOptionsDto) *dto.UserSingleRespDto {
	body, err := client.SendHttpRequest(
		"/api/v3/signup",
		fasthttp.MethodPost,
		&dto.SignUpDto{
			Connection: "PASSWORD",
			PasswordPayload: dto.SignUpByPasswordDto{
				Email:    email,
				Password: password,
			},
			Options: options,
		},
	)
	var response dto.UserSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

func (client *AuthenticationClient) SignUpByUsernamePassword(username string, password string, options dto.SignUpOptionsDto) *dto.UserSingleRespDto {
	body, err := client.SendHttpRequest(
		"/api/v3/signup",
		fasthttp.MethodPost,
		&dto.SignUpDto{
			Connection: "PASSWORD",
			PasswordPayload: dto.SignUpByPasswordDto{
				Username: username,
				Password: password,
			},
			Options: options,
		},
	)
	var response dto.UserSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

// ==== AUTO GENERATED AUTHENTICATION METHODS BEGIN ====
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
func (client *AuthenticationClient) SignInByCredentials(reqDto *dto.SigninByCredentialsDto) *dto.LoginTokenRespDto {
	b, err := client.SendHttpRequest("/api/v3/signin", fasthttp.MethodPost, reqDto)
	var response dto.LoginTokenRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) SignInByMobile(reqDto *dto.SigninByMobileDto) *dto.LoginTokenRespDto {
	b, err := client.SendHttpRequest("/api/v3/signin-by-mobile", fasthttp.MethodPost, reqDto)
	var response dto.LoginTokenRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) GetAlipayAuthInfo(reqDto *dto.GetAlipayAuthinfoDto) *dto.GetAlipayAuthInfoRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-alipay-authinfo", fasthttp.MethodGet, reqDto)
	var response dto.GetAlipayAuthInfoRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) GeneQrCode(reqDto *dto.GenerateQrcodeDto) *dto.GeneQRCodeRespDto {
	b, err := client.SendHttpRequest("/api/v3/gene-qrcode", fasthttp.MethodPost, reqDto)
	var response dto.GeneQRCodeRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) CheckQrCodeStatus(reqDto *dto.CheckQrcodeStatusDto) *dto.CheckQRCodeStatusRespDto {
	b, err := client.SendHttpRequest("/api/v3/check-qrcode-status", fasthttp.MethodGet, reqDto)
	var response dto.CheckQRCodeStatusRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) ExchangeTokenSetWithQrCodeTicket(reqDto *dto.ExchangeTokenSetWithQRcodeTicketDto) *dto.LoginTokenRespDto {
	b, err := client.SendHttpRequest("/api/v3/exchange-tokenset-with-qrcode-ticket", fasthttp.MethodPost, reqDto)
	var response dto.LoginTokenRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) ChangeQrCodeStatus(reqDto *dto.ChangeQRCodeStatusDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/change-qrcode-status", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) SendSms(reqDto *dto.SendSMSDto) *dto.SendSMSRespDto {
	b, err := client.SendHttpRequest("/api/v3/send-sms", fasthttp.MethodPost, reqDto)
	var response dto.SendSMSRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) SendEmail(reqDto *dto.SendEmailDto) *dto.SendEmailRespDto {
	b, err := client.SendHttpRequest("/api/v3/send-email", fasthttp.MethodPost, reqDto)
	var response dto.SendEmailRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) GetProfile(reqDto *dto.GetProfileDto) *dto.UserSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-profile", fasthttp.MethodGet, reqDto)
	var response dto.UserSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) UpdateProfile(reqDto *dto.UpdateUserProfileDto) *dto.UserSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-profile", fasthttp.MethodPost, reqDto)
	var response dto.UserSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) BindEmail(reqDto *dto.BindEmailDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/bind-email", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) UnbindEmail(reqDto *dto.UnbindEmailDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/unbind-email", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) BindPhone(reqDto *dto.BindPhoneDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/bind-phone", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) UnbindPhone(reqDto *dto.UnbindPhoneDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/unbind-phone", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) GetSecurityLevel() *dto.GetSecurityInfoRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-security-info", fasthttp.MethodGet, nil)
	var response dto.GetSecurityInfoRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) UpdatePassword(reqDto *dto.UpdatePasswordDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/update-password", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) VerifyUpdateEmailRequest(reqDto *dto.VerifyUpdateEmailRequestDto) *dto.VerifyUpdateEmailRequestRespDto {
	b, err := client.SendHttpRequest("/api/v3/verify-update-email-request", fasthttp.MethodPost, reqDto)
	var response dto.VerifyUpdateEmailRequestRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) UpdateEmail(reqDto *dto.UpdateEmailDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/update-email", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) VerifyUpdatePhoneRequest(reqDto *dto.VerifyUpdatePhoneRequestDto) *dto.VerifyUpdatePhoneRequestRespDto {
	b, err := client.SendHttpRequest("/api/v3/verify-update-phone-request", fasthttp.MethodPost, reqDto)
	var response dto.VerifyUpdatePhoneRequestRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) UpdatePhone(reqDto *dto.UpdatePhoneDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/update-phone", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) VerifyResetPasswordRequest(reqDto *dto.VerifyResetPasswordRequestDto) *dto.PasswordResetVerifyResp {
	b, err := client.SendHttpRequest("/api/v3/verify-reset-password-request", fasthttp.MethodPost, reqDto)
	var response dto.PasswordResetVerifyResp
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) ResetPassword(reqDto *dto.ResetPasswordDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/reset-password", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) VerifyDeleteAccountRequest(reqDto *dto.VerifyDeleteAccountRequestDto) *dto.VerifyDeleteAccountRequestRespDto {
	b, err := client.SendHttpRequest("/api/v3/verify-delete-account-request", fasthttp.MethodPost, reqDto)
	var response dto.VerifyDeleteAccountRequestRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) DeleteAccount(reqDto *dto.DeleteAccounDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-account", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) GetSystemInfo() *dto.SystemInfoResp {
	b, err := client.SendHttpRequest("/api/v3/system", fasthttp.MethodGet, nil)
	var response dto.SystemInfoResp
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) GetCountryList() *dto.GetCountryListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-country-list", fasthttp.MethodGet, nil)
	var response dto.GetCountryListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) PreCheckCode(reqDto *dto.PreCheckCodeDto) *dto.PreCheckCodeRespDto {
	b, err := client.SendHttpRequest("/api/v3/pre-check-code", fasthttp.MethodPost, reqDto)
	var response dto.PreCheckCodeRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
* @summary 发起绑定 MFA 认证要素请求
* @description 当用户未绑定某个 MFA 认证要素时，可以发起绑定 MFA 认证要素请求。不同类型的 MFA 认证要素绑定请求需要发送不同的参数，详细见 profile 参数。发起验证请求之后，Authing 服务器会根据相应的认证要素类型和传递的参数，使用不同的手段要求验证。此接口会返回 enrollmentToken，你需要在请求「绑定 MFA 认证要素」接口时带上此 enrollmentToken，并提供相应的凭证。
* @returns SendEnrollFactorRequestRespDto
 */
func (client *AuthenticationClient) SendEnrollFactorRequest(reqDto *dto.SendEnrollFactorRequestDto) *dto.SendEnrollFactorRequestRespDto {
	b, err := client.SendHttpRequest("/api/v3/send-enroll-factor-request", fasthttp.MethodPost, reqDto)
	var response dto.SendEnrollFactorRequestRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) EnrollFactor(reqDto *dto.EnrollFactorDto) *dto.EnrollFactorRespDto {
	b, err := client.SendHttpRequest("/api/v3/enroll-factor", fasthttp.MethodPost, reqDto)
	var response dto.EnrollFactorRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) ResetFactor(reqDto *dto.ResetFactorDto) *dto.ResetFactorRespDto {
	b, err := client.SendHttpRequest("/api/v3/reset-factor", fasthttp.MethodPost, reqDto)
	var response dto.ResetFactorRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) ListEnrolledFactors() *dto.ListEnrolledFactorsRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-enrolled-factors", fasthttp.MethodGet, nil)
	var response dto.ListEnrolledFactorsRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) GetFactor(reqDto *dto.GetFactorDto) *dto.GetFactorRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-factor", fasthttp.MethodGet, reqDto)
	var response dto.GetFactorRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) ListFactorsToEnroll() *dto.ListFactorsToEnrollRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-factors-to-enroll", fasthttp.MethodGet, nil)
	var response dto.ListFactorsToEnrollRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
	* @summary 生成绑定外部身份源的链接
	* @description
 * 此接口用于生成绑定外部身份源的链接，生成之后可以引导用户进行跳转。
 *
	* @returns GenerateBindExtIdpLinkRespDto
*/
func (client *AuthenticationClient) GenerateLinkExtIdpUrl(reqDto *dto.GenerateLinkExtidpUrlDto) *dto.GenerateBindExtIdpLinkRespDto {
	b, err := client.SendHttpRequest("/api/v3/generate-link-extidp-url", fasthttp.MethodGet, reqDto)
	var response dto.GenerateBindExtIdpLinkRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) UnlinkExtIdp(reqDto *dto.UnlinkExtIdpDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/unlink-extidp", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) GetIdentities() *dto.GetIdentitiesRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-identities", fasthttp.MethodGet, nil)
	var response dto.GetIdentitiesRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) GetApplicationEnabledExtIdps() *dto.GetExtIdpsRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-application-enabled-extidps", fasthttp.MethodGet, nil)
	var response dto.GetExtIdpsRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) SignUp(reqDto *dto.SignUpDto) *dto.UserSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/signup", fasthttp.MethodPost, reqDto)
	var response dto.UserSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) DecryptWechatMiniProgramData(reqDto *dto.DecryptWechatMiniProgramDataDto) *dto.DecryptWechatMiniProgramDataRespDto {
	b, err := client.SendHttpRequest("/api/v3/decrypt-wechat-miniprogram-data", fasthttp.MethodPost, reqDto)
	var response dto.DecryptWechatMiniProgramDataRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) GetWechatMiniprogramPhone(reqDto *dto.GetWechatMiniProgramPhoneDto) *dto.GetWechatMiniProgramPhoneRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-wechat-miniprogram-phone", fasthttp.MethodPost, reqDto)
	var response dto.GetWechatMiniProgramPhoneRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) GetWechatMpAccessToken(reqDto *dto.GetWechatAccessTokenDto) *dto.GetWechatAccessTokenRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-wechat-access-token", fasthttp.MethodPost, reqDto)
	var response dto.GetWechatAccessTokenRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) GetLoginHistory(reqDto *dto.GetMyLoginHistoryDto) *dto.GetLoginHistoryRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-my-login-history", fasthttp.MethodGet, reqDto)
	var response dto.GetLoginHistoryRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) GetLoggedInApps() *dto.GetLoggedInAppsRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-my-logged-in-apps", fasthttp.MethodGet, nil)
	var response dto.GetLoggedInAppsRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) GetAccessibleApps() *dto.GetAccessibleAppsRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-my-accessible-apps", fasthttp.MethodGet, nil)
	var response dto.GetAccessibleAppsRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
* @summary 获取租户列表
* @description 获取租户列表
* @returns GetTenantListRespDto
 */
func (client *AuthenticationClient) GetTenantList() *dto.GetTenantListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-my-tenant-list", fasthttp.MethodGet, nil)
	var response dto.GetTenantListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
* @summary 获取角色列表
* @description 获取角色列表
* @returns RoleListRespDto
 */
func (client *AuthenticationClient) GetRoleList(reqDto *dto.GetMyRoleListDto) *dto.RoleListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-my-role-list", fasthttp.MethodGet, reqDto)
	var response dto.RoleListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
* @summary 获取分组列表
* @description 获取分组列表
* @returns GroupListRespDto
 */
func (client *AuthenticationClient) GetGroupList() *dto.GroupListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-my-group-list", fasthttp.MethodGet, nil)
	var response dto.GroupListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) GetDepartmentList(reqDto *dto.GetMyDepartmentListDto) *dto.UserDepartmentPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-my-department-list", fasthttp.MethodGet, reqDto)
	var response dto.UserDepartmentPaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
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
func (client *AuthenticationClient) GetAuthorizedResources(reqDto *dto.GetMyAuthorizedResourcesDto) *dto.AuthorizedResourcePaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-my-authorized-resources", fasthttp.MethodGet, reqDto)
	var response dto.AuthorizedResourcePaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

// ==== AUTO GENERATED AUTHENTICATION METHODS END ====
