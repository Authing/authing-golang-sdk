package authentication

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"net/url"
)

type ContentType int

const (
	Default ContentType = iota // 开始生成枚举值, 默认为0
	Json
)

type ProtocolRequestOption struct {
	Url         string
	Method      string
	ReqDto      interface{}
	Headers     map[string]string
	ContentType ContentType
}

type ResponseData struct {
	Body       []byte
	Header     *fasthttp.ResponseHeader
	StatusCode int
}

func GenQueryString(variables map[string]interface{}) string {
	params := url.Values{}
	for key, value := range variables {
		params.Add(key, fmt.Sprintf("%v", value))
	}
	qs := params.Encode()
	return qs
}

func GenFormArgs(variables map[string]interface{}) *fasthttp.Args {
	args := &fasthttp.Args{}
	for key, value := range variables {
		args.Add(key, fmt.Sprintf("%v", value))
	}
	return args
}

func (client AuthenticationClient) SendProtocolHttpRequest(option *ProtocolRequestOption) (*ResponseData, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	method := option.Method
	reqDto := option.ReqDto
	url := option.Url
	if method == fasthttp.MethodGet && reqDto != nil {
		variables := reqDto.(map[string]interface{})
		qs := GenQueryString(variables)
		if qs != "" {
			url += "?" + qs
		}
	}

	req.SetRequestURI(url)
	for key, value := range option.Headers {
		req.Header.Add(key, value)
	}
	req.Header.SetMethod(method)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	if option.ContentType == Json {
		req.Header.SetContentType("application/json; charset=UTF-8")
		bytes, err := json.Marshal(reqDto) //序列化json

		if err != nil {
			return nil, err
		}
		req.SetBody(bytes)

	} else if method == fasthttp.MethodPost {
		req.Header.SetContentType("application/x-www-form-urlencoded; charset=UTF-8")
		if reqDto != nil {
			variables := reqDto.(map[string]interface{})
			bytes := GenFormArgs(variables).QueryString()
			req.SetBody(bytes)
		}
	} else if method == fasthttp.MethodGet {

	} else {
		return nil, fmt.Errorf("不支持的请求类型")
	}

	var httpClient = &fasthttp.Client{
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	err := httpClient.Do(req, resp)
	if err != nil {
		return nil, err
	}
	statusCode := resp.StatusCode()
	body := resp.Body()

	header := fasthttp.ResponseHeader{}
	resp.Header.CopyTo(&header)
	return &ResponseData{
		Body:       body,
		Header:     &header,
		StatusCode: statusCode,
	}, nil
}
