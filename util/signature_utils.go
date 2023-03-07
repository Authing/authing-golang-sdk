package util

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"sort"
	"strings"
)

func sortForeach(data map[string]string, action func(string, string)) {
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		action(k, data[k])
	}
}

func signature(secret, stringToSign string) string {
	key := []byte(secret)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func ComposeStringToSign(method, uri string, headers, queries map[string]string) string {
	stringToSign := method + "\n"
	headerLen := len(headers)
	if headerLen > 0 {
		sortForeach(headers, func(s1, s2 string) {
			stringToSign += s1 + ":" + s2 + "\n"
		})
	}
	uriParts := strings.Split(uri, "?")
	stringToSign += uriParts[0]
	if len(uriParts) > 1 {
		queries[uriParts[1]] = ""
	}
	if len(queries) > 0 {
		stringToSign += "?"
		sortForeach(queries, func(s1, s2 string) {
			stringToSign += s1
			if s2 != "" {
				stringToSign += "=" + s2
			}
			stringToSign += "&"
		})
		stringToSign = stringToSign[:len(stringToSign)-1]
	}
	return stringToSign
}

func GetAuthorization(accessKeyId, accessKeySecret, stringToSign string) string {
	signature := signature(accessKeySecret, stringToSign)
	return "authing " + accessKeyId + ":" + signature
}
