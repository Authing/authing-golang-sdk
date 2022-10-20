package management

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/Authing/authing-golang-sdk/constant"
	"github.com/Authing/authing-golang-sdk/dto"
	"github.com/Authing/authing-golang-sdk/util/cache"
	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
	"strings"
	"sync"
	"time"
)

type JwtClaims struct {
	*jwt.StandardClaims
	//用户编号
	UID      string
	Username string
}

func GetAccessToken(client *ManagementClient) (string, error) {
	// 从缓存获取token
	cacheToken, b := cache.GetCache(constant.TokenCacheKeyPrefix + client.options.AccessKeyId)
	if b && cacheToken != nil {
		return cacheToken.(string), nil
	}
	// 从服务获取token，加锁
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	cacheToken, b = cache.GetCache(constant.TokenCacheKeyPrefix + client.options.AccessKeyId)
	if b && cacheToken != nil {
		return cacheToken.(string), nil
	}
	resp, err := QueryAccessToken(client)
	if err != nil {
		return "", err
	}

	if token, _ := jwt.Parse(resp.Data.AccessToken, nil); token != nil {
		userPoolId := token.Claims.(jwt.MapClaims)["scoped_userpool_id"]
		client.userPoolId = userPoolId.(string)
	}

	cache.SetCache(constant.TokenCacheKeyPrefix+client.options.AccessKeyId, resp.Data.AccessToken, time.Duration(resp.Data.ExpiresIn*int(time.Second)))
	return resp.Data.AccessToken, nil
}

func QueryAccessToken(client *ManagementClient) (*dto.GetManagementTokenRespDto, error) {
	variables := map[string]interface{}{
		"accessKeyId":     client.options.AccessKeyId,
		"accessKeySecret": client.options.AccessKeySecret,
	}

	b, err := client.SendHttpRequest("/api/v3/get-management-token", fasthttp.MethodPost, variables)
	if err != nil {
		return nil, err
	}
	var r dto.GetManagementTokenRespDto
	if b != nil {
		json.Unmarshal(b, &r)
	}
	return &r, nil
}

func (c *ManagementClient) SendHttpRequest(url string, method string, reqDto interface{}) ([]byte, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(reqDto)
	if err != nil {
		return nil, err
	}
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

	req.SetRequestURI(c.options.Host + url)

	if method != fasthttp.MethodGet {
		req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	}
	req.Header.Add("x-authing-app-tenant-id", ""+c.options.TenantId)
	//req.Header.Add("x-authing-request-from", c.options.RequestFrom)
	req.Header.Add("x-authing-sdk-version", constant.SdkVersion)
	req.Header.Add("x-authing-lang", c.options.Lang)
	if url != "/api/v3/get-management-token" {
		token, _ := GetAccessToken(c)
		req.Header.Add("Authorization", "Bearer "+token)
		req.Header.Add("x-authing-userpool-id", c.userPoolId)
	}
	req.Header.SetMethod(method)

	bytes, err := json.Marshal(reqDto) //data是请求数据

	if err != nil {
		return nil, err
	}
	req.SetBody(bytes)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	client := &fasthttp.Client{
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client.Do(req, resp)
	body := resp.Body()
	return body, err
}