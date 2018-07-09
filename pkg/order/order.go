package order

import (
	"encoding/json"
	"time"
)

const (
	resourceURL = "orders"
	// Order Type
	LimitType  = "LIMIT"
	MarketType = "MARKET"
	// Order Side
	SellSide = "sell"
	BuySide  = "buy"
)

type baseOrder struct {
	ID        int64
	Symbol    string
	Market    string
	Price     json.Number
	Qty       json.Number
	RemainQty json.Number
	Cost      json.Number
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Side      int
	Type      int
}

type Order struct {
	baseOrder
	Side string
	Type string
}

func (o *Order) UnmarshalJSON(b []byte) error {
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

	return nil
}
