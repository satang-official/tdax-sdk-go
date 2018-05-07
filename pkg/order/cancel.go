package order

import (
	"github.com/satang-official/tdax-sdk-go/pkg/client"
	"github.com/satang-official/tdax-sdk-go/pkg/signature"
	resty "gopkg.in/resty.v1"
)

func Cancel(c client.Client, orderID int64) (*resty.Response, error) {
	sig := signature.Sign(c.APISecret(), nil)

	return resty.R().
		SetHeader("Authorization", c.APIKey()).
		SetHeader("Signature", sig).
		Delete(c.URL() + resourceURL + "/" + string(orderID))
}
