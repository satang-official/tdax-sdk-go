package withdrawal

import (
	"encoding/json"

	"github.com/satang-official/tdax-sdk-go/pkg/client"
	"github.com/satang-official/tdax-sdk-go/pkg/signature"
	resty "gopkg.in/resty.v1"
)

type wdFields struct {
	Address string
	Amount  string
	Tag     string `json:"omitempty"`
}

type Option struct {
	Type   string
	Fields wdFields
}

func NewWithdrawal(wdType, addr, amount, tag string) Option {
	return Option{
		Type: wdType,
		Fields: wdFields{
			Address: addr,
			Amount:  amount,
			Tag:     tag,
		},
	}
}

func Create(c client.Client, option Option) (*resty.Response, error) {
	bb, err := json.Marshal(option)
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{}
	json.Unmarshal(bb, &params)

	sig := signature.Sign(c.APISecret(), params)

	return resty.R().
		SetHeader("Authorization", c.APIKey()).
		SetHeader("Signature", sig).
		SetBody(option).
		Post(c.URL() + resourceURL)
}
