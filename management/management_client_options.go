package management

import (
	"time"

	"github.com/Authing/authing-golang-sdk/v3/constant"
	"github.com/Authing/authing-golang-sdk/v3/util"
	"github.com/valyala/fasthttp"
)

type ManagementClient struct {
	httpClient *fasthttp.Client
	options    *ManagementClientOptions
	userPoolId string
	eventHub   *util.WebSocketEventHub
}

type ManagementClientOptions struct {
	AccessKeyId     string
	AccessKeySecret string
	TenantId        string
	//Deprecated: Use ReadTimeout instead
	Timeout     int
	ReadTimeout time.Duration
	Lang        string
	Host        string
	/**
	是否跳过 HTTPS 证书检测，默认为 false；如果是私有化部署的场景且证书不被信任，可以设置为 true
	*/
	InsecureSkipVerify bool
	WssHost            string
	/**
	 * 自定义 Client 创建函数
	 */
	CreateClientFunc func(options *ManagementClientOptions) *fasthttp.Client
}

func NewManagementClient(options *ManagementClientOptions) (*ManagementClient, error) {
	if options.Host == "" {
		options.Host = constant.ApiServiceUrl
	}
	if options.WssHost == "" {
		options.WssHost = constant.WebSocketHost
	}
	if options.ReadTimeout == 0 {
		options.ReadTimeout = 10 * time.Second
	}
	c := &ManagementClient{
		options: options,
	}

	c.httpClient = c.createHttpClient()
	c.eventHub = util.NewWebSocketEvent()
	return c, nil
}
