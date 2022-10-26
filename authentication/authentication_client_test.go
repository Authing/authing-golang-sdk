package authentication

import (
	"fmt"
	"github.com/Authing/authing-golang-sdk/v3/constant"
	"github.com/Authing/authing-golang-sdk/v3/dto"

	// "fmt"
	"strings"
	"testing"

	"github.com/valyala/fasthttp"
)

var authenticationClient *AuthenticationClient
var options = AuthenticationClientOptions{
	AppId:              "",
	AppSecret:          "",
	AppHost:            "",
	RedirectUri:        "http://localhost:8989",
	InsecureSkipVerify: true,
}

const idToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiO" +
	"iI2MjkwNzU3ODliNDI0M2E1MGY2YzA0NTYiLCJiaXJ0aGRhdGUiOm51bGwsImZhbWls" +
	"eV9uYW1lIjpudWxsLCJnZW5kZXIiOiJVIiwiZ2l2ZW5fbmFtZSI6bnVsbCwibG9jYWx" +
	"lIjpudWxsLCJtaWRkbGVfbmFtZSI6bnVsbCwibmFtZSI6bnVsbCwibmlja25hbWUiOm" +
	"51bGwsInBpY3R1cmUiOiJodHRwczovL2ZpbGVzLmF1dGhpbmcuY28vYXV0aGluZy1jb" +
	"25zb2xlL2RlZmF1bHQtdXNlci1hdmF0YXIucG5nIiwicHJlZmVycmVkX3VzZXJuYW1l" +
	"IjpudWxsLCJwcm9maWxlIjpudWxsLCJ1cGRhdGVkX2F0IjoiMjAyMi0wNi0xNFQxMjo" +
	"zNDoxNC4yNjFaIiwid2Vic2l0ZSI6bnVsbCwiem9uZWluZm8iOm51bGwsImF0X2hhc2" +
	"giOiJGdkctUFBDWks5YlFMQ0dtRUxYMkZ3IiwiYXVkIjoiNjJhODU3MGE4NTg1OWUyM" +
	"zkwZWYzODhmIiwiZXhwIjoxNjU2NTA4MzQ1LCJpYXQiOjE2NTUyOTg3NDUsImlzcyI6" +
	"Imh0dHBzOi8vbG9jYWx0ZXN0LnRlc3QyLmF1dGhpbmctaW5jLmNvL29pZGMifQ.X50a" +
	"K3EYwmhUcRtgOKJXFIDYRr8gfaRDukg8Ev697Wo"

func init() {

	var err error
	authenticationClient, err = NewAuthenticationClient(&options)
	if err != nil {
		panic(err)
	}
}

func getReqUrl(path string) string {
	return "https://" + options.AppHost + path
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
	result, err := authenticationClient.BuildAuthorizeUrlByOidc(&OIDCAuthURLParams{
		Scope: "offline_access " + constant.DefaultScope,
	})
	if err != nil {
		panic(err)
		return
	}
	println(result.Url)
	//fmt.Println(result)
	//res, err1 := util.SendRequest(&util.RequestOption{
	//	Url:    result.Url,
	//	Method: fasthttp.MethodGet,
	//})
	//if err1 != nil {
	//	t.Fatalf("请求授权url失败 %v", err1)
	//	return
	//}
	//if res.StatusCode >= 400 {
	//	t.Fatalf("请求授权时失败：[%d]", res.StatusCode)
	//	return
	//}
	//loginPath := res.Header.Peek("Location")
	//urlEle := strings.Split(string(loginPath), "?")
	//if len(urlEle) < 2 {
	//	t.Fatalf("授权地址格式错误 %s", loginPath)
	//	return
	//}
	//pathEle := strings.Split(urlEle[0], "/")
	//uuid := pathEle[len(pathEle)-1]
	//fmt.Println(uuid)
	//cookieStr := getCookieStr(res.Header)
	//headers := map[string]string{
	//	"cookie": cookieStr,
	//}
	//res, err1 = util.SendRequest(&util.RequestOption{
	//	Url:     getReqUrl(string(loginPath)),
	//	Headers: headers,
	//	Method:  fasthttp.MethodGet,
	//})
	//if err1 != nil {
	//	t.Fatalf("请求登录url失败 %v", err1)
	//	return
	//}
	//if res.StatusCode >= 400 {
	//	t.Fatalf("请求登录url时失败： %d", res.StatusCode)
	//	return
	//}
}

func TestCode(t *testing.T) {
	tokenResponse, err := authenticationClient.GetAccessTokenByCode("BP7D0_o3Ya0TudEP3VolHVVywFDo_e3DFm-19koxQwy")
	if err != nil {
		t.Fatalf("code校验失败, %v", err)
		return
	}
	fmt.Println(tokenResponse)
}

func TestIntrospectToken(t *testing.T) {
	code := "e1I4h2L-9-BaaL87YKtZgjKxsUHppaPW2jsLCsEokuL"
	tokenResponse, _ := authenticationClient.GetAccessTokenByCode(code)
	resp1, _ := authenticationClient.IntrospectAccessTokenOffline(tokenResponse.AccessToken)
	fmt.Printf("%+v\n", resp1.Id)
	resp2, _ := authenticationClient.IntrospectAccessTokenOffline(tokenResponse.RefreshToken)
	resp3, _ := authenticationClient.IntrospectAccessTokenOffline(tokenResponse.IDToken)
	fmt.Printf("%+v\n", resp1)
	fmt.Printf("%+v\n", resp2)
	fmt.Printf("%+v\n", resp3)
	result1, _ := authenticationClient.IntrospectToken(tokenResponse.AccessToken)
	result2, _ := authenticationClient.IntrospectToken(tokenResponse.RefreshToken)
	result3, _ := authenticationClient.IntrospectToken(tokenResponse.IDToken)
	fmt.Printf("%+v\n", result1)
	fmt.Printf("%+v\n", result2)
	fmt.Printf("%+v\n", result3)
	result4, _ := authenticationClient.RevokeToken(tokenResponse.AccessToken)
	fmt.Printf("%+v\n", result4)
}

func TestIDToken(t *testing.T) {

	charim, err := authenticationClient.ParseIDToken(idToken)
	if err != nil {
		t.Fatalf("id token 校验失败, %v", err)
		return
	}
	fmt.Println(charim.IssuedAt)
}

//func TestRreshToken(t *testing.T) {
//	refreshToken := "XbOJEYqDkKh71taxISZO-ICxQexljlTmXQGf6dZNVOs"
//	tokens, err := authenticationClient.RefreshLoginState(refreshToken)
//	if err != nil {
//		t.Fatalf("测试刷新token失败: %v", err)
//		return
//	}
//	fmt.Println(tokens)
//}

func TestLogout(t *testing.T) {
	url, err := authenticationClient.BuildLogoutUrl(&BuildLogoutURLParams{
		IDTokenHint: idToken,
	})
	if err != nil {
		t.Fatalf("生成退出url失败:%v", err)
	}
	fmt.Println(url)
}

func Test_SignInByUsernamePassword(t *testing.T) {
	resp := authenticationClient.SignInByEmailPassword(
		"test@example.com", "test", dto.SignInOptionsDto{})
	fmt.Println(resp.StatusCode, resp.RequestId, resp.Message)
	fmt.Println(resp.Data.AccessToken)
	authenticationClient.SetAccessToken(resp.Data.AccessToken)

	profileResp, err := authenticationClient.RevokeToken(resp.Data.AccessToken)
	fmt.Println(profileResp, err)

	//str, err := authenticationClient.BuildLogoutUrl(resp.Data.AccessToken)
	//fmt.Println(str, err)
}
