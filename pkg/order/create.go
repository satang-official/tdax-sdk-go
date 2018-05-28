package order

import (
	"encoding/json"
	"math/big"
	"time"

	"github.com/satang-official/tdax-sdk-go/pkg/client"
	"github.com/satang-official/tdax-sdk-go/pkg/signature"
	resty "gopkg.in/resty.v1"
)

type Option struct {
	Type   string
	Symbol string
	Market string
	Price  string `json:",omitempty"`
	Qty    string `json:",omitempty"`
	Side   string
	Nonce  int64
}

// NewLimitSell is Order construct helper
func NewLimitSell(price, qty *big.Int, symbol, market string) Option {
	return Option{
		Type:   LimitType,
		Side:   SellSide,
		Nonce:  time.Now().Unix(),
		Price:  price.String(),
		Qty:    qty.String(),
		Symbol: symbol,
		Market: market,
	}
}

// NewLimitBuy is Order construct helper
func NewLimitBuy(price, qty *big.Int, symbol, market string) Option {
	return Option{
		Type:   LimitType,
		Side:   BuySide,
		Nonce:  time.Now().Unix(),
		Price:  price.String(),
		Qty:    qty.String(),
		Symbol: symbol,
		Market: market,
	}
}

// NewMarketSell is Order construct helper
func NewMarketSell(price, qty *big.Int, symbol, market string) Option {
	return Option{
		Type:   LimitType,
		Side:   SellSide,
		Nonce:  time.Now().Unix(),
		Price:  "",
		Qty:    qty.String(),
		Symbol: symbol,
		Market: market,
	}
}

// NewMarketBuy is Order construct helper
func NewMarketBuy(price, qty *big.Int, symbol, market string) Option {
	return Option{
		Type:   MarketType,
		Side:   BuySide,
		Nonce:  time.Now().Unix(),
		Price:  price.String(),
		Qty:    "",
		Symbol: symbol,
		Market: market,
	}
}

func Create(c client.Client, option Option) (*Order, error) {
	bb, err := json.Marshal(option)
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{}
	json.Unmarshal(bb, &params)

	sig := signature.Sign(c.APISecret(), params)

	order := &Order{}

	_, err = resty.R().
		SetHeader("Authorization", c.APIKey()).
		SetHeader("Signature", sig).
		SetBody(option).
		SetResult(order).
		Post(c.URL() + resourceURL)
	return order, err
}
