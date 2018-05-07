package order

import (
	"encoding/json"
	"math/big"
	"time"

	"bitbucket.org/satangcorp/tdax-sdk/pkg/client"
	"bitbucket.org/satangcorp/tdax-sdk/pkg/signature"
	resty "gopkg.in/resty.v1"
)

type Option struct {
	Type   string
	Symbol string
	Market string
	Price  string
	Qty    string
	Side   string
	Nonce  int64
}

const limit = "LIMIT"
const sell = "sell"
const buy = "buy"

func NewLimitSell(price, qty *big.Int, symbol, market string) Option {
	return Option{
		Type:   limit,
		Side:   sell,
		Nonce:  time.Now().Unix(),
		Price:  price.String(),
		Qty:    qty.String(),
		Symbol: symbol,
		Market: market,
	}
}

func NewLimitBuy(price, qty *big.Int, symbol, market string) Option {
	return Option{
		Type:   limit,
		Side:   buy,
		Nonce:  time.Now().Unix(),
		Price:  price.String(),
		Qty:    qty.String(),
		Symbol: symbol,
		Market: market,
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
