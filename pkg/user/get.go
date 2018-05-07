package user

import (
	"github.com/satang-official/tdax-sdk-go/pkg/client"
	"github.com/satang-official/tdax-sdk-go/pkg/signature"
	resty "gopkg.in/resty.v1"
)

func Get(c client.Client) (*resty.Response, error) {
	sig := signature.Sign(c.APISecret(), nil)
	return resty.R().
		SetHeader("Authorization", c.APIKey()).
		SetHeader("Signature", sig).
		Get(c.URL() + "/users/" + c.UserID())
}
