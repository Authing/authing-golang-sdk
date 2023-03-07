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
	AppId:       "",
	AppSecret:   "",
	AppHost:     "",
	WssHost:     "ws://localhost:88",
	RedirectUri: "http://localhost:8989",
	AccessToken: "",
}

const idToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiI2M2IzZWE3Mzk0MDc0YWE1ZTE5YmIzMGMiLCJhdWQiOiI2M2IzZWE1YmM5NjY0YzlkMDhkMTkwYzgiLCJpYXQiOjE2NzI3MzY2NjcsImV4cCI6MTY3Mzk0NjI2NywiaXNzIjoiaHR0cHM6Ly9nby11c2VyLXBlcm1pc3Npb24tYXV0aC5hdXRoaW5nLmNuL29pZGMiLCJub25jZSI6InJtZ0pab3RxNGkiLCJuYW1lIjpudWxsLCJnaXZlbl9uYW1lIjpudWxsLCJtaWRkbGVfbmFtZSI6bnVsbCwiZmFtaWx5X25hbWUiOm51bGwsIm5pY2tuYW1lIjpudWxsLCJwcmVmZXJyZWRfdXNlcm5hbWUiOm51bGwsInByb2ZpbGUiOm51bGwsInBpY3R1cmUiOiJodHRwczovL2ZpbGVzLmF1dGhpbmcuY28vYXV0aGluZy1jb25zb2xlL2RlZmF1bHQtdXNlci1hdmF0YXIucG5nIiwid2Vic2l0ZSI6bnVsbCwiYmlydGhkYXRlIjpudWxsLCJnZW5kZXIiOiJVIiwiem9uZWluZm8iOm51bGwsImxvY2FsZSI6bnVsbCwidXBkYXRlZF9hdCI6IjIwMjMtMDEtMDNUMDg6NTQ6NDUuMzMwWiIsImVtYWlsIjoidGVzdEBleGFtcGxlLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwicGhvbmVfbnVtYmVyIjpudWxsLCJwaG9uZV9udW1iZXJfdmVyaWZpZWQiOmZhbHNlfQ.3awp567aJ3wBXR0mh0l1oBTugTNsDqYpJVIaDTeHXbI\n"

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

//func TestIntrospectToken(t *testing.T) {
//	code := "e1I4h2L-9-BaaL87YKtZgjKxsUHppaPW2jsLCsEokuL"
//	tokenResponse, _ := authenticationClient.GetAccessTokenByCode(code)
//	resp1, _ := authenticationClient.IntrospectAccessTokenOffline(tokenResponse.AccessToken)
//	fmt.Printf("%+v\n", resp1.Id)
//	resp2, _ := authenticationClient.IntrospectAccessTokenOffline(tokenResponse.RefreshToken)
//	resp3, _ := authenticationClient.IntrospectAccessTokenOffline(tokenResponse.IDToken)
//	fmt.Printf("%+v\n", resp1)
//	fmt.Printf("%+v\n", resp2)
//	fmt.Printf("%+v\n", resp3)
//	result1, _ := authenticationClient.IntrospectToken(tokenResponse.AccessToken)
//	result2, _ := authenticationClient.IntrospectToken(tokenResponse.RefreshToken)
//	result3, _ := authenticationClient.IntrospectToken(tokenResponse.IDToken)
//	fmt.Printf("%+v\n", result1)
//	fmt.Printf("%+v\n", result2)
//	fmt.Printf("%+v\n", result3)
//	result4, _ := authenticationClient.RevokeToken(tokenResponse.AccessToken)
//	fmt.Printf("%+v\n", result4)
//}

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

func TestAuthenticationClient_CheckPermissionByStringResource(t *testing.T) {

	request := dto.CheckPermissionStringResourceDto{
		Resources: []string{"stringResourceCode1", "stringResourceCode2"},
		Action:    "delete",
	}
	response := authenticationClient.CheckPermissionByStringResource(&request)
	fmt.Println(response)

}

func TestAuthenticationClient_CheckPermissionByArrayResource(t *testing.T) {

	resp := authenticationClient.SignInByEmailPassword(
		"test@example.com", "test", dto.SignInOptionsDto{})
	fmt.Println(resp.StatusCode, resp.RequestId, resp.Message)
	fmt.Println(resp.Data.AccessToken)
	authenticationClient.SetAccessToken(resp.Data.AccessToken)

	request := dto.CheckPermissionArrayResourceDto{
		Resources: []string{"arrayResourceCode1", "arrayResourceCode2"},
		Action:    "get",
	}
	response := authenticationClient.CheckPermissionByArrayResource(&request)
	fmt.Println(response)

}

func TestAuthenticationClient_CheckPermissionByTreeResource(t *testing.T) {

	resp := authenticationClient.SignInByEmailPassword(
		"test@example.com", "test", dto.SignInOptionsDto{})
	fmt.Println(resp.StatusCode, resp.RequestId, resp.Message)
	fmt.Println(resp.Data.AccessToken)
	authenticationClient.SetAccessToken(resp.Data.AccessToken)

	request := dto.CheckPermissionTreeResourceDto{
		Resources: []string{"treeResourceCode/tree1", "treeResourceCode/tree1/tree11", "treeResourceCode/tree1/tree11/tree111", "treeResourceCode/tree2", "treeResourceCode/tree2/tree22", "treeResourceCode/tree2/tree22/tree222"},
		Action:    "get",
	}
	response := authenticationClient.CheckPermissionByTreeResource(&request)
	fmt.Println(response)

}

func TestAuthenticationClient_GetUserAuthorizedResourcesList(t *testing.T) {

	resp := authenticationClient.SignInByEmailPassword(
		"test@example.com", "test", dto.SignInOptionsDto{})
	fmt.Println(resp.StatusCode, resp.RequestId, resp.Message)
	fmt.Println(resp.Data.AccessToken)
	authenticationClient.SetAccessToken(resp.Data.AccessToken)

	response := authenticationClient.GetUserAuthorizedResourcesList()
	fmt.Println(response)

}

func TestAuthenticationClient_GetUserAuthResourcePermissionList(t *testing.T) {

	resp := authenticationClient.SignInByUsernamePassword(
		"4", "4", dto.SignInOptionsDto{})
	fmt.Println(resp.StatusCode, resp.RequestId, resp.Message)
	fmt.Println(resp.Data.AccessToken)
	authenticationClient.SetAccessToken(resp.Data.AccessToken)

	request := dto.GetUserAuthResourcePermissionListDto{
		Resources: []string{"r1", "r2", "r3/1", "r3/1/11", "r3/1/11/111", "r3/2"},
	}
	response := authenticationClient.getUserAuthResourcePermissionList(&request)
	fmt.Println(response)

}

func TestAuthenticationClient_GetUserAuthResourceStruct(t *testing.T) {

	resp := authenticationClient.SignInByUsernamePassword(
		"4", "4", dto.SignInOptionsDto{})
	fmt.Println(resp.StatusCode, resp.RequestId, resp.Message)
	fmt.Println(resp.Data.AccessToken)
	authenticationClient.SetAccessToken(resp.Data.AccessToken)

	request := dto.GetUserAuthResourceStructDto{
		Resource: "r3",
	}

	response := authenticationClient.getUserAuthResourceStruct(&request)
	fmt.Println(response)

}

type UserEvent struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func TestClient_PubEvent(t *testing.T) {
	var data = &UserEvent{
		Id:   "232323232",
		Name: "golang-authentication-test",
	}
	var resp = authenticationClient.PubEvent("custom_xrgu.user_create", data)
	fmt.Println(resp)
}

type Receiver1 struct{}

func (receiver *Receiver1) OnSuccess(msg []byte) {
	fmt.Println(string(msg))
}
func (receiver *Receiver1) OnError(err error) {
	fmt.Println(err)
}

func TestClient_SubEvent(t *testing.T) {
	ErrChan := make(chan error, 1)
	receiver := &Receiver1{}
	authenticationClient.SubEventByReceiver("custom_xrgu.user_create", receiver)

	authenticationClient.SubEvent("custom_xrgu.user_create", func(msg []byte) {
		fmt.Println(string(msg) + "222")
	}, func(err error) {
		fmt.Println(err)
		ErrChan <- err
	})

	authenticationClient.SubEvent("custom_xrgu.user_create", func(msg []byte) {
		fmt.Println(string(msg) + "333")
	}, func(err error) {
		fmt.Println(err)
		ErrChan <- err
	})
	<-ErrChan
}
