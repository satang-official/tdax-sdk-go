package order

import (
	"fmt"

	"github.com/satang-official/tdax-sdk-go/pkg/client"
	"github.com/satang-official/tdax-sdk-go/pkg/signature"
	resty "gopkg.in/resty.v1"
)

type ListResp struct {
	Count int
	Items []Order
}

func List(c client.Client, params ...string) (*ListResp, error) {
	sig := signature.Sign(c.APISecret(), nil)

	pair := ""
	if len(params) == 2 {
		pair = fmt.Sprintf("%s_%s", params[0], params[1])
	}

	orders := ListResp{Count: -1}
	req := resty.R().
		SetHeader("Authorization", c.APIKey()).
		SetHeader("Signature", sig).
		SetQueryParams(map[string]string{
			"Symbol": pair,
			"Limit":  "50",
			"Offset": "0",
			"Status": "open",
		})

	if pair == "" {
		req.SetResult(&orders)
	} else {
		req.SetResult(&orders.Items)
	}

	_, err := req.Get(c.URL() + resourceURL)
	if err != nil {
		return nil, err
	}
	return &orders, err
}
