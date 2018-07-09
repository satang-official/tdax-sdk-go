package order

import (
	"github.com/satang-official/tdax-sdk-go/pkg/client"
	"gopkg.in/resty.v1"
	"github.com/satang-official/tdax-sdk-go/pkg/signature"
	"strconv"
	"encoding/json"
)

type Status string

var status = []string{ "SUBMITTED", "REJECTED", "ACCEPTED", "PARTIAL_FILLED", "CANCELLED","FILLED"}

type StatusOrder struct {
	baseOrder
	Status string
	Side string
	Type string
}

func (o *StatusOrder) UnmarshalJSON(b []byte) error {
	tmpO := baseOrder{}
	err := json.Unmarshal(b, &tmpO)
	if err != nil {
		return err
	}

	o.baseOrder = tmpO
	if tmpO.Side == 0 {
		o.Side = BuySide
	} else {
		o.Side = SellSide
	}
	if tmpO.Type == 0 {
		o.Type = LimitType
	} else {
		o.Type = MarketType
	}
	if tmpO.Status < len(status) {
		o.Status = status[tmpO.Status]
	}

	return nil
}

func Get(c client.Client, id int64) (StatusOrder, error) {
	sig := signature.Sign(c.APISecret(), nil)

	order := StatusOrder{}
	req := resty.R().
		SetHeader("Authorization", c.APIKey()).
		SetHeader("Signature", sig).
		SetResult(&order)

	_, err := req.Get(c.URL() + resourceURL + "/" + strconv.FormatInt(id, 10))
	return order, err

}
