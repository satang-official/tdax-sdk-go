package order

import (
	"fmt"

	"bitbucket.org/satangcorp/tdax-sdk/pkg/client"
	"bitbucket.org/satangcorp/tdax-sdk/pkg/signature"
	resty "gopkg.in/resty.v1"
)

const resourceURL = "orders"

func List(c client.Client, symbol, market string) (*resty.Response, error) {
	sig := signature.Sign(c.APISecret(), nil)

	pair := fmt.Sprintf("%s_%s", symbol, market)

	return resty.R().
		SetHeader("Authorization", c.APIKey()).
		SetHeader("Signature", sig).
		SetQueryParams(map[string]string{
			"Symbol": pair,
			"Limit":  "20",
			"Offset": "0",
			"Status": "open",
		}).
		Get(c.URL() + resourceURL)
}
