package client

import (
	"encoding/json"
	"errors"
	"fmt"

	// "net/http"
	"strings"
	"time"

	"authing-go-sdk/constant"
	"authing-go-sdk/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
	"github.com/MicahParks/compatibility-keyfunc"
)
const RandStringLen = 16
const ALG_HS256 = "HS256"
const JWK_PATH = "/oidc/.well-known/jwks.json"
type AuthenticationClient struct {
	appId string
	appSecret string
	domain string
	redirectUri string;
	logoutRedirectUri string;
	scope string
	jwks *keyfunc.JWKS
}

func NewAuthenticationClient(options *AuthenticationClientOptions) (*AuthenticationClient, error) {
	if options.AppId == "" {
		return nil, errors.New("AppId不能为空")
	}
	if options.AppSecret == "" {
		return nil, errors.New("AppSecret不能为空")
	}
	if options.Domain == "" {
		return nil, errors.New("Domain不能为空")
	}
	if options.RedirectUri == "" {
		return nil, errors.New("RedirectUri不能为空")
	}
	client := &AuthenticationClient {
		appId: options.AppId,
		appSecret: options.AppSecret,
		domain: options.Domain,
		redirectUri: options.RedirectUri,
		logoutRedirectUri: options.LogoutRedirectUri,
	};
	if options.Scope == "" {
		client.scope = constant.DefaultScope;
	}

	jwksURL := client.getUrl(JWK_PATH)
	jwks, err := keyfunc.Get(jwksURL, keyfunc.Options{})
	if err != nil {
		return nil, err
	}
	client.jwks = jwks

	return client, nil;
}

func (client *AuthenticationClient) getUrl(path string) string {
	return "https://" + client.domain + path
}

func (client *AuthenticationClient) BuildAuthUrl(params *AuthURLParams) (AuthUrlResult, error) {
	scope := params.Scope;
	if scope == "" {
		scope = client.scope
	}
	state := params.State
	if state == "" {
		;// none to do
	}
	nonce := params.Nonce
	if nonce == "" {
		nonce = util.RandStringImpr(RandStringLen)
	}
	redirectUri := params.RedirectUri
	if redirectUri == "" {
		redirectUri = client.redirectUri
	}
	dataMap := map[string]interface{} {
		"redirect_uri": redirectUri,
		"client_id": client.appId,
		"response_mode": "query",
		"response_type": "code",
		"scope": scope,
		"nonce": nonce,
	}
	if state != "" {
		dataMap["state"] = state
	}
	if (params.Forced) {
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
	return AuthUrlResult {
		State: state,
		Nonce: nonce,
		Url: client.getUrl("/oidc/auth?") + util.GenQueryString(dataMap),
	}, nil
}

func (client *AuthenticationClient) GetLoginStateByAuthCode(params *CodeToTokenParams) (*LoginState, error) {
	redirectUrl := params.RedirectUri;
	if redirectUrl == "" {
		redirectUrl = client.redirectUri
	}
	data := map[string]interface{} {
		"code": params.Code,
		"client_id": client.appId,
		"client_secret": client.appSecret,
		"grant_type": "authorization_code",
		"redirect_uri": redirectUrl,
	}
	res, err := util.SendRequest(&util.RequestOption{
		Url: client.getUrl("/oidc/token"),
		Method: fasthttp.MethodPost,
		ReqDto: data,
	})
	if err != nil {
		return nil, err
	}
	if (res.StatusCode >= 400) {
		return nil, fmt.Errorf("根据code获取token失败[%d]:%s", res.StatusCode, res.Body)
	}
	// dataStr := string(res.Header.Peek("Date"))
	// serverTime, _ := http.ParseTime(dataStr) 
	fmt.Println(string(res.Body))
	return client.buildLoginState(res.Body)
}

func (client *AuthenticationClient) getKeyCommon(token *jwt.Token) (interface{}, error) {
	alg, ok := token.Header["alg"].(string)
	if !ok {
		return nil, fmt.Errorf("算法字段非法 %v", token.Header["alg"])
	}
	if alg == ALG_HS256 {
		return []byte(client.appSecret), nil
	}
	return client.jwks.KeyfuncLegacy(token)
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

func (client *AuthenticationClient) ParsedIDToken(tokenStr string) (*IDTokenClaims, error){
	tokenJwt, err := jwt.ParseWithClaims(tokenStr, &IDTokenClaims{}, client.getKey4IdToken)
	if err != nil {
		return nil, err
	}
	claims, ok := tokenJwt.Claims.(*IDTokenClaims)
	if ok && tokenJwt.Valid {
		return claims, nil
	}
	// TODO return error
	return nil, nil	
}

func (client *AuthenticationClient) ParsedAccessToken(tokenStr string) (*AccessTokenClaims, error){
	token, err := jwt.ParseWithClaims(tokenStr, &AccessTokenClaims{},  client.getKey4AccessToken)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*AccessTokenClaims)
	if ok && token.Valid {
		return claims, nil
	}
	// TODO return error
	return nil, nil
}

func (client *AuthenticationClient) buildLoginState(bytes []byte ) (*LoginState, error) {
	var loginState LoginState
	err := json.Unmarshal(bytes, &loginState)
	if err != nil {
		return nil, err
	}
	var second time.Duration = time.Duration(loginState.ExpiresIn) * time.Second
	var date time.Time = time.Now()
	loginState.ExpireAt = date.Add(second)
	loginState.ParsedIDToken, err = client.ParsedIDToken(loginState.IdToken)
	if err != nil {
		return nil, err
	}
	loginState.ParsedAccessToken, err = client.ParsedAccessToken(loginState.AccessToken)
	if err != nil {
		return nil, err
	}
	return &loginState, nil
}

