package authentication

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Authing/authing-golang-sdk/v3/constant"
	"github.com/Authing/authing-golang-sdk/v3/util"
	"github.com/valyala/fasthttp"
)

func (client *AuthenticationClient) SendHttpRequest(url string, method string, reqDto interface{}) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	data, _ := json.Marshal(&reqDto)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)

	var queryString strings.Builder
	if method == fasthttp.MethodGet {
		if variables != nil && len(variables) > 0 {
			for key, value := range variables {
				queryString.WriteString(key)
				queryString.WriteString("=")
				queryString.WriteString(fmt.Sprintf("%v", value))
				queryString.WriteString("&")
			}
		}
		qs := queryString.String()
		if qs != "" {
			url += "?" + qs
		}
	}

	// 设置请求地址
	req.SetRequestURI(client.options.AppHost + url)

	// 设置请求头
	if method != fasthttp.MethodGet {
		req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	}
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

	// 设置请求方法
	req.Header.SetMethod(method)

	bytes, err := json.Marshal(reqDto) //data是请求数据

	if err != nil {
		return nil, err
	}
	req.SetBody(bytes)
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
