package xsig

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"sort"

	"github.com/sirupsen/logrus"
)

func ConstructQueryString(params map[string]interface{}) string {
	return constructQueryStringWithPrefix(params, "")
}

func Sign(secret []byte, params map[string]interface{}) []byte {
	qs := ConstructQueryString(params)

	mac := hmac.New(sha512.New, secret)
	mac.Write([]byte(qs))
	logrus.Errorf("foo %+v", qs)

	return mac.Sum(nil)
}

func Verify(secret []byte, params map[string]interface{}, signature []byte) bool {
	return hmac.Equal(Sign(secret, params), signature)
}

func constructQueryStringWithPrefix(params map[string]interface{}, prefix string) string {
	var keys []string

	for k, _ := range params {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	var qs string

	for i, k := range keys {
		v := params[k]

		if nestedParams, ok := v.(map[string]interface{}); ok {
			qs += constructQueryStringWithPrefix(nestedParams, k)
		} else if array, ok := v.([]interface{}); ok {
			logrus.Errorf("bar")
			nestedMap := map[string]interface{}{}
			for i, v := range array {
				nestedMap[fmt.Sprintf("%d", i)] = v
			}
			qs += constructQueryStringWithPrefix(nestedMap, k)
		} else {
			if prefix == "" {
				if _, ok := v.(float64); ok {
					qs += fmt.Sprintf("%s=%.f", k, v)
				} else {
					qs += fmt.Sprintf("%s=%v", k, v)
				}
			} else {
				if _, ok := v.(float64); ok {
					qs += fmt.Sprintf("%s[%s]=%.f", prefix, k, v)
				} else {
					qs += fmt.Sprintf("%s[%s]=%v", prefix, k, v)
				}
			}
		}

		if i != len(keys)-1 {
			qs += "&"
		}
	}

	return qs
}
