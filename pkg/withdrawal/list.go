package withdrawal

import (
	"github.com/satang-official/tdax-sdk-go/pkg/client"
	"github.com/satang-official/tdax-sdk-go/pkg/signature"
	resty "gopkg.in/resty.v1"
)

const resourceURL = "withdrawals"

func List(c client.Client) (*resty.Response, error) {
	sig := signature.Sign(c.APISecret(), nil)

	return resty.R().
		SetHeader("Authorization", c.APIKey()).
		SetHeader("Signature", sig).
		Get(c.URL() + resourceURL)
}
