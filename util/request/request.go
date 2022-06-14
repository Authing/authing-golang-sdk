package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/valyala/fasthttp"
)

func SendRequest(url string, method string, reqDto interface{}, headers map[string]string) ([]byte, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(reqDto)
	if err != nil {
		return nil, err
	}
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	var queryString strings.Builder
	if method == fasthttp.MethodGet {
		variables := reqDto.(map[string]interface{})
		for key, value := range variables {
			queryString.WriteString(key)
			queryString.WriteString("=")
			queryString.WriteString(fmt.Sprintf("%v", value))
			queryString.WriteString("&")
		}
		qs := queryString.String()
		if qs != "" {
			url += "?" + qs
		}
	}

	req.SetRequestURI(url)

	if method != fasthttp.MethodGet {
		req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	req.Header.SetMethod(method)

	bytes, err := json.Marshal(reqDto) //data是请求数据

	if err != nil {
		return nil, err
	}
	req.SetBody(bytes)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	client := &fasthttp.Client{}
	client.Do(req, resp)
	body := resp.Body()
	return body, err
}
