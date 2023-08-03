package util

import (
	"fmt"
	"net/url"

	"github.com/Authing/authing-golang-sdk/v3/constant"
)

func GetQueryString(queryMap map[string]string) string {
	if queryMap == nil || len(queryMap) == 0 {
		return constant.StringEmpty
	}
	queryValue := url.Values{}
	for key, value := range queryMap {
		if value == "" {
			continue
		}
		queryValue.Add(key, value)
	}
	return queryValue.Encode()
}

func GetQueryString2(queryMap map[string]interface{}) string {
	if queryMap == nil || len(queryMap) == 0 {
		return constant.StringEmpty
	}
	queryValue := url.Values{}
	for key, value := range queryMap {
		if value == "" {
			continue
		}
		queryValue.Add(key, fmt.Sprint(value))
	}
	return queryValue.Encode()
}
