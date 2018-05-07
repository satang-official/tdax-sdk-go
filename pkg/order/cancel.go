package order

import (
	"bitbucket.org/satangcorp/tdax-sdk/pkg/client"
	"bitbucket.org/satangcorp/tdax-sdk/pkg/signature"
	resty "gopkg.in/resty.v1"
)

func Cancel(c client.Client, orderID int64) (*resty.Response, error) {
	sig := signature.Sign(c.APISecret(), nil)

	return resty.R().
		SetHeader("Authorization", c.APIKey()).
		SetHeader("Signature", sig).
		Delete(c.URL() + resourceURL + "/" + string(orderID))
}
