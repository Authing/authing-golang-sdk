package management

import (
	"crypto/tls"
	"github.com/Authing/authing-golang-sdk/constant"
	"net/http"
)

type ManagementClient struct {
	HttpClient *http.Client
	options    *ClientOptions
	userPoolId string
}

type ClientOptions struct {
	AccessKeyId     string
	AccessKeySecret string
	TenantId        string
	Timeout         int
	Lang            string
	Host            string
	/**
	是否拒绝非法的 HTTPS 请求，默认为 true；如果是私有化部署的场景且证书不被信任，可以设置为 false
	*/
	RejectUnauthorized bool
}

func NewClient(options *ClientOptions) (*ManagementClient, error) {
	if options.Host == "" {
		options.Host = constant.ApiServiceUrl
	}
	c := &ManagementClient{
		options: options,
	}
	if c.HttpClient == nil {
		c.HttpClient = &http.Client{}
		_, err := GetAccessToken(c)
		if err != nil {
			return nil, err
		}
		/*src := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: accessToken},
		)
		c.HttpClient = oauth2.NewClient(context.Background(), src)*/
	}

	if c.options.RejectUnauthorized == false {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	return c, nil
}
