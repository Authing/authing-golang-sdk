package authentication

import (
	"encoding/json"
	"fmt"

	"github.com/Authing/authing-golang-sdk/v3/util"
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

func (client AuthenticationClient) SendProtocolHttpRequest(option *ProtocolRequestOption) (*ResponseData, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	method := option.Method
	reqDto := option.ReqDto
	url := option.Url
	queryString := util.GetQueryString(reqDto)
	if method == fasthttp.MethodGet && queryString != "" {
		url += "?" + queryString
	}

	req.Header.SetMethod(method)
	req.SetRequestURI(url)

	for key, value := range option.Headers {
		req.Header.Add(key, value)
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	switch method {
	case fasthttp.MethodPost:
		if option.ContentType == Json {
			req.Header.SetContentType("application/json; charset=UTF-8")
			bytes, err := json.Marshal(reqDto) //序列化json
			if err != nil {
				return &ResponseData{
					StatusCode: 500,
				}, err
			}
			req.SetBody(bytes)
		} else {
			req.Header.SetContentType("application/x-www-form-urlencoded; charset=UTF-8")
			if queryString != "" {
				req.SetBodyString(queryString)
			}
		}
	case fasthttp.MethodGet:
	default:
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
