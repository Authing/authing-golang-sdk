package client

import (
	"authing-go-sdk/util"
	"fmt"
	// "fmt"
	"strings"
	"testing"

	"github.com/valyala/fasthttp"
)

var clientAuth *AuthenticationClient
var options = AuthenticationClientOptions{
	AppId:     "62a8570a85859e2390ef388f",
	AppSecret: "ffe0ecad57823426e065a8c6d6bcd0b8",
	Domain: "localtest.test2.authing-inc.co",
	RedirectUri: "http://localhost:7001/callback",
}
func init() {
	
	var err error
	clientAuth, err = NewAuthenticationClient(&options)
	if err != nil {
		panic(err)
	}
}

func getReqUrl(path string) string {
	return "https://" + options.Domain + path
}

func getCookieStr(header *fasthttp.ResponseHeader) string {
	var cookieStr strings.Builder
	header.VisitAllCookie(func(key, value []byte) {
		cookieStr.WriteString(string(key))
		cookieStr.WriteString("=")
		cookie := fasthttp.Cookie{}
		cookie.ParseBytes(value)
		cookieStr.WriteString(string(cookie.Value()))
		cookieStr.WriteString("; ")
	})
	return strings.Trim(cookieStr.String(), "; ")
}

func TestAuthUrl(t *testing.T) {
	result, err := clientAuth.BuildAuthUrl(&AuthURLParams{})
	if err != nil {
		t.Fatalf("构建授权url失败 %v", err)
		return;
	}
	fmt.Println(result)
	res, err1 := util.SendRequest(&util.RequestOption{
		Url: result.Url,
		Method: fasthttp.MethodGet,
	})
	if err1 != nil {
		t.Fatalf("请求授权url失败 %v", err1)
		return;
	}
	if res.StatusCode >= 400 {
		t.Fatalf("请求授权时失败：[%d]", res.StatusCode)
		return
	}
	loginPath := res.Header.Peek("Location")
	urlEle := strings.Split(string(loginPath), "?")
	if len(urlEle) < 2 {
		t.Fatalf("授权地址格式错误 %s", loginPath)
		return
	}
	pathEle := strings.Split(urlEle[0], "/")
	uuid := pathEle[len(pathEle) - 1]
	fmt.Println(uuid)
	cookieStr := getCookieStr(res.Header)
	headers := map[string]string {
		"cookie": cookieStr,
	}
	res, err1 = util.SendRequest(&util.RequestOption{
		Url: getReqUrl(string(loginPath)),
		Headers:headers,
		Method: fasthttp.MethodGet,
	})
	if err1 != nil {
		t.Fatalf("请求登录url失败 %v", err1)
		return;
	}
	if res.StatusCode >= 400 {
		t.Fatalf("请求登录url时失败： %d", res.StatusCode)
		return;
	}

}

func TestCode(t *testing.T) {
	loginState, err := clientAuth.GetLoginStateByAuthCode(&CodeToTokenParams{
		Code: "vg5eVYlE60dLCxFzarmjB5p5Ede2lY0dJmGfinJnBqg",
	})
	if err != nil {
		t.Fatalf("code校验失败, %v", err)
		return
	}
	fmt.Println(loginState)
}