package util

import (
	"fmt"
	"net/url"
)

func GenQueryString(variables map[string]interface{}) string {
	params := url.Values{}
	for key, value := range variables {
		params.Add(key, fmt.Sprintf("%v", value))
	}
	qs := params.Encode()
	return qs
}
