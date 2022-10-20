package authentication

import (
	"fmt"
	"github.com/Authing/authing-golang-sdk/constant"
	"github.com/Authing/authing-golang-sdk/dto"

	// "fmt"
	"strings"
	"testing"

	"github.com/valyala/fasthttp"
)

var authenticationClient *AuthenticationClient
var options = AuthenticationClientOptions{
	AppId:              "6343b99bd968375153712c3c",
	AppSecret:          "2140f44a44fcf7684954e37c25d70ce4",
	AppHost:            "https://sadsdd.cj.mereith.com",
	RedirectUri:        "http://localhost:3003/callback",
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
		t.Fatalf("构建授权url失败 %v", err)
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

//func TestCode(t *testing.T) {
//	loginState, err := authenticationClient.GetLoginStateByAuthCode(&CodeToTokenParams{
//		Code: "g1FZq2O8y3NzHvn3YwtTW7dau6lJD9Icq2ZTUR88d_a",
//	})
//	if err != nil {
//		t.Fatalf("code校验失败, %v", err)
//		return
//	}
//	fmt.Println(loginState)
//}

func TestAccessToken(t *testing.T) {
	accessToken := `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6I` +
		`jROVS13OVZIcmVjU1BuT20zNzJubVF4V0ROV1hQbUQxbDdBckNseXhyVTAifQ.eyJ` +
		`qdGkiOiJWeFdMRVJPY0FjSy0xR240Y0M3UGciLCJzdWIiOiI2MjkwNzU3ODliNDI0` +
		`M2E1MGY2YzA0NTYiLCJpYXQiOjE2NTUyOTgxNzMsImV4cCI6MTY1NjUwNzc3Mywic` +
		`2NvcGUiOiJvcGVuaWQgcHJvZmlsZSIsImlzcyI6Imh0dHBzOi8vbG9jYWx0ZXN0Ln` +
		`Rlc3QyLmF1dGhpbmctaW5jLmNvL29pZGMiLCJhdWQiOiI2MmE4NTcwYTg1ODU5ZTI` +
		`zOTBlZjM4OGYifQ.c64QBODEI_u1KQJaTi_00kz-zquXBwndwvKSRRc2N0LQBX9Ki` +
		`mObyLBLEodkdZH61k-JVtI1IFlyupYB1QxejyxpfsbKMCokJ7JaM4J9l1I4Sre9RZ` +
		`5CFrP3I03p0eEGiPSfLx3zBswfTz__b9ClnxyAGy3vqj69j3BZxK139ocnG39LHqg` +
		`svZ5thY8w4iwFqZE3lZwKNPRdbaRnC5YyP6Y9M8xP9sQNiRTNxNGZPazCsj1RZWhK` +
		`VP8a71QyTydSPccIi6s4-GzusO5iKC2bPEGtjwYaWlIK_C-cJtGhXwoYppbUP5sQV` +
		`tVUPTVtbua_KYomBjsVIoGaeadV-cg1TA`
	charim, err := authenticationClient.IntrospectAccessTokenOffline(accessToken)
	if err != nil {
		t.Fatalf("access token 校验失败, %v", err)
		return
	}
	fmt.Println(charim)
	user, err1 := authenticationClient.GetUserInfo(accessToken)
	if err1 != nil {
		t.Fatalf("获取用户信息失败，%v", err1)
	}
	fmt.Println(user)
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
	url, err := authenticationClient.BuildLogoutUrl(&LogoutURLParams{
		IDTokenHint: idToken,
	})
	if err != nil {
		t.Fatalf("生成退出url失败:%v", err)
	}
	fmt.Println(url)
}

func Test_SignInByUsernamePassword(t *testing.T) {
	resp := authenticationClient.SignInByUsernamePassword(
		"test", "test", dto.SignInOptionsDto{})
	println(resp.StatusCode, resp.RequestId, resp.Message)
	println(resp.Data.AccessToken)
	authenticationClient.SetAccessToken(resp.Data.AccessToken)

	profileResp := authenticationClient.GetProfile(&dto.GetProfileDto{})
	println(profileResp.Data.UserId)
}
