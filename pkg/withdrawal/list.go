package withdrawal

import (
	"bitbucket.org/satangcorp/tdax-sdk/pkg/client"
	"bitbucket.org/satangcorp/tdax-sdk/pkg/signature"
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
