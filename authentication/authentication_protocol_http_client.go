package authentication

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/valyala/fasthttp"
)

type ContentType int

const (
	Default ContentType = iota // 开始生成枚举值, 默认为0
	Json
)

type ProtocolRequestOption struct {
	Url         string
	Method      string
	ReqDto      map[string]string
	Headers     map[string]string
	ContentType ContentType
}

type ResponseData struct {
	Body       []byte
	Header     *fasthttp.ResponseHeader
	StatusCode int
}

func GenQueryString(variables map[string]string) string {
	params := url.Values{}
	for key, value := range variables {
		params.Add(key, fmt.Sprintf("%v", value))
	}
	qs := params.Encode()
	return qs
}

func GenFormArgs(variables map[string]string) *fasthttp.Args {
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
		variables := reqDto
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
			return &ResponseData{
				StatusCode: 500,
			}, err
		}
		req.SetBody(bytes)

	} else if method == fasthttp.MethodPost {
		req.Header.SetContentType("application/x-www-form-urlencoded; charset=UTF-8")
		if reqDto != nil {
			variables := reqDto
			bytes := GenFormArgs(variables).QueryString()
			req.SetBody(bytes)
		}
	} else if method == fasthttp.MethodGet {

	} else {
		return &ResponseData{
			StatusCode: 500,
		}, fmt.Errorf("不支持的请求类型")
	}

	err := client.httpClient.DoTimeout(req, resp, client.options.ReadTimeout)
	if err != nil {
		return &ResponseData{
			StatusCode: 500,
		}, err
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
