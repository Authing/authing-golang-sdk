package authentication

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/Authing/authing-golang-sdk/v3/constant"
	"github.com/Authing/authing-golang-sdk/v3/util"
	"github.com/valyala/fasthttp"
)

func (client *AuthenticationClient) SendHttpRequest(url string, method string, reqDto interface{}) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	reqJsonBytes, err := json.Marshal(&reqDto)
	if err != nil {
		return nil, err
	}
	if method == fasthttp.MethodGet {
		variables := make(map[string]interface{})
		err = json.Unmarshal(reqJsonBytes, &variables)
		if err != nil {
			return nil, err
		}
		queryString := util.GetQueryString2(variables)
		if queryString != "" {
			url += "?" + queryString
		}
	}

	// 设置请求方法
	req.Header.SetMethod(method)
	// 设置请求地址
	req.SetRequestURI(client.options.AppHost + url)

	//req.Header.Add("x-authing-request-from", c.options.RequestFrom)
	req.Header.Add("x-authing-sdk-version", constant.SdkVersion)
	//req.Header.Add("x-authing-lang", c.Lang)
	req.Header.Add("x-authing-app-id", client.options.AppId)

	endpointsToSendBasicHeader := []string{
		"/oidc/token",
		"/oidc/token/revocation",
		"/oidc/token/introspection",
		"/oauth/token",
		"/oauth/token/revocation",
		"/oauth/token/introspection",
		"/api/v3/signin",
		"/api/v3/signin-by-mobile",
		"/api/v3/exchange-tokenset-with-qrcode-ticket",
	}
	if client.options.TokenEndPointAuthMethod == ClientSecretBasic && util.StringContains(endpointsToSendBasicHeader, url) {
		req.Header.Add("authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", client.options.AppId, client.options.AppSecret))))
	} else if client.options.AccessToken != "" {
		req.Header.Add("authorization", client.options.AccessToken)
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")

	if method != fasthttp.MethodGet {
		req.SetBody(reqJsonBytes)
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err = client.httpClient.DoTimeout(req, resp, client.options.ReadTimeout)
	if err != nil {
		resultMap := make(map[string]interface{})
		if err == fasthttp.ErrTimeout {
			resultMap["statusCode"] = 504
			resultMap["message"] = "请求超时"
		} else {
			resultMap["statusCode"] = 500
			resultMap["message"] = err.Error()
		}
		b, err := json.Marshal(resultMap)
		if err != nil {
			return nil, err
		}
		return b, err
	}
	body := resp.Body()
	return body, err
}

func (client *AuthenticationClient) createHttpClient() *fasthttp.Client {
	options := client.options
	createClientFunc := options.CreateClientFunc
	if createClientFunc != nil {
		return createClientFunc(options)
	}
	return &fasthttp.Client{
		TLSConfig: &tls.Config{InsecureSkipVerify: options.InsecureSkipVerify},
	}
}
