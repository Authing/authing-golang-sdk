package management

import (
	"github.com/Authing/authing-golang-sdk/constant"
	"net/http"
)

type ManagementClient struct {
	HttpClient *http.Client
	options    *ManagementClientOptions
	userPoolId string
}

type ManagementClientOptions struct {
	AccessKeyId     string
	AccessKeySecret string
	TenantId        string
	Timeout         int
	Lang            string
	Host            string
	/**
	是否跳过 HTTPS 证书检测，默认为 false；如果是私有化部署的场景且证书不被信任，可以设置为 true
	*/
	InsecureSkipVerify bool
}

func NewManagementClient(options *ManagementClientOptions) (*ManagementClient, error) {
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
		c.HttpClient = oauth2.NewManagementClient(context.Background(), src)*/
	}
	return c, nil
}
