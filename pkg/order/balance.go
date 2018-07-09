package order

import (
	"github.com/satang-official/tdax-sdk-go/pkg/signature"
	"github.com/satang-official/tdax-sdk-go/pkg/client"
	"gopkg.in/resty.v1"
)

type Balance struct {
	Buy map[string]float64
	Sell map[string]float64
}

func GetBalance(c client.Client) (Balance, error) {
	sig := signature.Sign(c.APISecret(), nil)

	balance := Balance{}
	req := resty.R().
		SetHeader("Authorization", c.APIKey()).
		SetHeader("Signature", sig).
		SetResult(&balance)

	_, err := req.Get(c.URL() + resourceURL + "/balances")
	return balance, err
}

