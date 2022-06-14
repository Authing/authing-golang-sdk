package client

import (
	"errors"

	"authing-go-sdk/constant"
)

type AuthenticationClient struct {
	appId string
	appSecret string
	domain string
	redirectUri string;
	logoutRedirectUri string;
	scope string;
	// cookieKey string;
}

func NewAuthenticationClient(options *AuthenticationClientOptions) (*AuthenticationClient, error) {
	if options.AppId == "" {
		return nil, errors.New("AppId不能为空")
	}
	if options.AppSecret == "" {
		return nil, errors.New("AppSecret不能为空")
	}
	if options.Domain == "" {
		return nil, errors.New("Domain不能为空")
	}
	if options.RedirectUri == "" {
		return nil, errors.New("RedirectUri不能为空")
	}
	client := &AuthenticationClient {
		appId: options.AppId,
		appSecret: options.AppSecret,
		domain: options.Domain,
		redirectUri: options.RedirectUri,
		logoutRedirectUri: options.LogoutRedirectUri,
	};
	if options.Scope == "" {
		client.scope = constant.DefaultScope;
	}

	return client, nil;
}

func (client *AuthenticationClient) BuildAuthorizeUrl() (string, error) {
	return "", nil
} 