package management

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	url2 "net/url"
	"sync"
	"time"

	"github.com/Authing/authing-golang-sdk/v3/constant"
	"github.com/Authing/authing-golang-sdk/v3/dto"
	"github.com/Authing/authing-golang-sdk/v3/util/cache"
	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
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
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	data, _ := json.Marshal(&reqDto)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)

	if method == fasthttp.MethodGet && variables != nil && len(variables) > 0 {
		params := url2.Values{}
		urlTemp, _ := url2.Parse(url)
		for key, value := range variables {
			params.Set(key, fmt.Sprintf("%v", value))
		}
		urlTemp.RawQuery = params.Encode()
		url = urlTemp.String()
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

	err = client.DoTimeout(req, resp, c.options.ReadTimeout)
	if err != nil {
		if err == fasthttp.ErrTimeout {
			resultMap := make(map[string]interface{})
			resultMap["statusCode"] = 504
			resultMap["message"] = "请求超时"
			b, err := json.Marshal(resultMap)
			if err != nil {
				return nil, err
			}
			return b, err
		}
		return nil, err
	}
	body := resp.Body()
	return body, err
}
