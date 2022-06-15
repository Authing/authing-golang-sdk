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
		Code: "fTXlN2TMkcqZNVB94HS-ulah1mZdBgBdEQHN4k1AFd9",
	})
	if err != nil {
		t.Fatalf("code校验失败, %v", err)
		return
	}
	fmt.Println(loginState)
}

func TestAccessToken(t *testing.T) {
	charim, err := clientAuth.ParsedAccessToken("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IjROVS13OVZIcmVjU1BuT20zNzJubVF4V0ROV1hQbUQxbDdBckNseXhyVTAifQ.eyJqdGkiOiJWeFdMRVJPY0FjSy0xR240Y0M3UGciLCJzdWIiOiI2MjkwNzU3ODliNDI0M2E1MGY2YzA0NTYiLCJpYXQiOjE2NTUyOTgxNzMsImV4cCI6MTY1NjUwNzc3Mywic2NvcGUiOiJvcGVuaWQgcHJvZmlsZSIsImlzcyI6Imh0dHBzOi8vbG9jYWx0ZXN0LnRlc3QyLmF1dGhpbmctaW5jLmNvL29pZGMiLCJhdWQiOiI2MmE4NTcwYTg1ODU5ZTIzOTBlZjM4OGYifQ.c64QBODEI_u1KQJaTi_00kz-zquXBwndwvKSRRc2N0LQBX9KimObyLBLEodkdZH61k-JVtI1IFlyupYB1QxejyxpfsbKMCokJ7JaM4J9l1I4Sre9RZ5CFrP3I03p0eEGiPSfLx3zBswfTz__b9ClnxyAGy3vqj69j3BZxK139ocnG39LHqgsvZ5thY8w4iwFqZE3lZwKNPRdbaRnC5YyP6Y9M8xP9sQNiRTNxNGZPazCsj1RZWhKVP8a71QyTydSPccIi6s4-GzusO5iKC2bPEGtjwYaWlIK_C-cJtGhXwoYppbUP5sQVtVUPTVtbua_KYomBjsVIoGaeadV-cg1TA")
	if err != nil {
		t.Fatalf("access token 校验失败, %v", err)
		return
	}
	fmt.Println(charim)
}

func TestIDToken(t *testing.T) {
	charim, err := clientAuth.ParsedIDToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI2MjkwNzU3ODliNDI0M2E1MGY2YzA0NTYiLCJiaXJ0aGRhdGUiOm51bGwsImZhbWlseV9uYW1lIjpudWxsLCJnZW5kZXIiOiJVIiwiZ2l2ZW5fbmFtZSI6bnVsbCwibG9jYWxlIjpudWxsLCJtaWRkbGVfbmFtZSI6bnVsbCwibmFtZSI6bnVsbCwibmlja25hbWUiOm51bGwsInBpY3R1cmUiOiJodHRwczovL2ZpbGVzLmF1dGhpbmcuY28vYXV0aGluZy1jb25zb2xlL2RlZmF1bHQtdXNlci1hdmF0YXIucG5nIiwicHJlZmVycmVkX3VzZXJuYW1lIjpudWxsLCJwcm9maWxlIjpudWxsLCJ1cGRhdGVkX2F0IjoiMjAyMi0wNi0xNFQxMjozNDoxNC4yNjFaIiwid2Vic2l0ZSI6bnVsbCwiem9uZWluZm8iOm51bGwsImF0X2hhc2giOiJGdkctUFBDWks5YlFMQ0dtRUxYMkZ3IiwiYXVkIjoiNjJhODU3MGE4NTg1OWUyMzkwZWYzODhmIiwiZXhwIjoxNjU2NTA4MzQ1LCJpYXQiOjE2NTUyOTg3NDUsImlzcyI6Imh0dHBzOi8vbG9jYWx0ZXN0LnRlc3QyLmF1dGhpbmctaW5jLmNvL29pZGMifQ.X50aK3EYwmhUcRtgOKJXFIDYRr8gfaRDukg8Ev697Wo")
	if err != nil {
		t.Fatalf("id token 校验失败, %v", err)
		return
	}
	fmt.Println(charim.IssuedAt)
}